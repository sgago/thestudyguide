package replicas

import (
	"fmt"
	"net/http"
	"net/url"
	"sgago/thestudyguide-causal/broadcast"
	"sgago/thestudyguide-causal/deduplicator"
	"sgago/thestudyguide-causal/lamport"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	causalPath  = "/causal"
	healthzPath = "/healthz"
	incrPath    = "/increment"
	doIncrPath  = "/do-increment"

	causalHealthzPath = causalPath + healthzPath
	causalIncrPath    = causalPath + incrPath
	causalDoIncrPath  = causalPath + doIncrPath
)

var (
	counter atomic.Uint64
)

type Replicas struct {
	router *gin.Engine
	client resty.Client
	dedupe *deduplicator.Deduplicator[uint64]
	clock  *lamport.Clock

	hosts []string
}

func New(router *gin.Engine, resty *resty.Client, clock *lamport.Clock, names ...string) *Replicas {
	hosts := make([]string, 0, len(names))

	for _, other := range names {
		hosts = append(hosts, fmt.Sprintf("http://%s:%d", other, 8080))
	}

	replicas := &Replicas{
		router: router,
		client: *resty,
		hosts:  hosts,
		dedupe: deduplicator.New[uint64](1 * time.Minute),
		clock:  clock,
	}

	replicas.addRoutes(replicas.router)

	return replicas
}

func (r *Replicas) addRoutes(router *gin.Engine) {
	group := router.Group(causalPath)
	{
		group.GET(healthzPath, func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
		})

		group.POST(incrPath, func(c *gin.Context) {
			<-r.BroadcastIncrement()
		})

		group.POST(doIncrPath, r.Increment)
	}
}

func (r *Replicas) urls(path ...string) []string {
	var urls []string

	for _, host := range r.hosts {
		u, _ := url.JoinPath(host, path...)
		urls = append(urls, u)
	}

	return urls
}

func (r *Replicas) BroadcastIncrement() chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := r.client.R().
			SetHeader("X-Deduplication-Id", strconv.FormatUint(counter.Load(), 10)).
			SetHeader("X-Replica-Id", "1")

		urls := r.urls(causalDoIncrPath)

		result := <-broadcast.
			Post(req, urls...).
			OnSuccess(func(resp *resty.Response) {
				fmt.Println(string(resp.Body()))
				fmt.Println(resp.Request.GenerateCurlCommand())
			}).
			OnError(func(err error) {
				fmt.Println(err)
			}).
			Send()

		done <- result
	}()

	return done
}

func (r *Replicas) Increment(c *gin.Context) {
	header := c.GetHeader("X-Deduplication-Id")
	dedupeId, err := strconv.ParseUint(header, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid deduplication id",
		})
		return
	}

	if r.dedupe.Mark(dedupeId) {
		counter.Add(1)

		<-r.BroadcastIncrement()
	}
}

func (r *Replicas) Healthz() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := r.client.R()
		urls := r.urls(causalHealthzPath)

		result := <-broadcast.
			Post(req, urls...).
			OnSuccess(func(resp *resty.Response) {
				fmt.Println(resp.Request.GenerateCurlCommand())
			}).
			OnError(func(err error) {
				fmt.Println(err)
			}).
			Send()

		done <- result
	}()

	return done
}
