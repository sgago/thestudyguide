package replicas

import (
	"fmt"
	"net/http"
	"net/url"
	"sgago/thestudyguide-causal/broadcast"
	"sgago/thestudyguide-causal/deduplicator"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	causal  = "/causal"
	healthz = "/healthz"

	causalHealthz = causal + healthz
)

var (
	counter atomic.Uint64
)

type Replicas struct {
	router *gin.Engine
	client resty.Client
	dedupe *deduplicator.Deduplicator[uint64]

	hosts []string
}

func New(router *gin.Engine, resty *resty.Client, names ...string) *Replicas {
	hosts := make([]string, 0, len(names))

	for _, other := range names {
		hosts = append(hosts, fmt.Sprintf("http://%s:%d", other, 8080))
	}

	replicas := &Replicas{
		router: router,
		client: *resty,
		hosts:  hosts,
		dedupe: deduplicator.New[uint64](1 * time.Minute),
	}

	replicas.addRoutes(replicas.router)

	return replicas
}

func (r *Replicas) addRoutes(router *gin.Engine) {
	group := router.Group(causal)
	{
		group.GET(healthz, func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
		})

		group.GET("/count", func(c *gin.Context) {

		})
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

func (r *Replicas) BroadcastIncrement() {

}

func (r *Replicas) Increment() {

}

func (r *Replicas) Healthz() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := r.client.R()
		urls := r.urls(causalHealthz)

		done <- <-broadcast.
			Get(req, urls...).
			OnSuccess(func(resp *resty.Response) {
				fmt.Println(resp.Request.GenerateCurlCommand())
			}).
			OnError(func(err error) {
				fmt.Println(err)
			}).
			Send()
	}()

	return done
}
