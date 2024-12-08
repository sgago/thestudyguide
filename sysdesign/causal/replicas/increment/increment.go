package increment

import (
	"fmt"
	"net/http"
	"sgago/thestudyguide-causal/broadcast"
	"sgago/thestudyguide-causal/deduplicator"
	"sgago/thestudyguide-causal/lamport"
	"sgago/thestudyguide-causal/replicas"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	incrPath = "/increment"

	Path = replicas.GroupPath + incrPath
)

type Increment struct {
	*replicas.Replicas

	counter atomic.Uint64
	dedupe  *deduplicator.Deduplicator[uint64]
	clock   *lamport.Clock
}

func New(r *replicas.Replicas, c *lamport.Clock) *Increment {
	incr := &Increment{
		Replicas: r,
		dedupe:   deduplicator.New[uint64](1 * time.Minute),
		clock:    c,
	}

	r.Router.POST(Path, incr.handle)

	return incr
}

func (incr *Increment) Do() <-chan bool {
	return incr.one(incr.clock.Get())
}

func (incr *Increment) broadcast() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		// TODO: Use consts, make more consistent
		req := incr.
			Request().
			SetHeader("X-Deduplication-Id", incr.clock.String()). // TODO: wrong, use lamport clock and incr correctly
			SetHeader("X-Replica-Id", strconv.Itoa(incr.MyId))

		urls := incr.Urls(Path)

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

func (incr *Increment) handle(c *gin.Context) {
	dedupeHeader := c.GetHeader("X-Deduplication-Id")
	dedupeId, err := strconv.ParseUint(dedupeHeader, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid deduplication id %s", dedupeHeader),
		})
		return
	}

	replicaIdHeader := c.GetHeader("X-Replica-Id")
	_, err = strconv.Atoi(replicaIdHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid replica id",
		})
		return
	}

	<-incr.one(dedupeId)
}

func (incr *Increment) one(dedupeId uint64) <-chan bool {
	done := make(<-chan bool, 1)

	if incr.dedupe.Mark(dedupeId) {
		incr.clock.Set(dedupeId)
		done = incr.broadcast() // Start the async broadcast
		incr.counter.Add(1)     // Do work in parallel
	}

	return done
}
