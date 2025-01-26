package counter

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
	Path     = replicas.GroupPath + incrPath
)

type Counter struct {
	*replicas.Replicas

	hostname string
	counter  atomic.Uint64
	dedupe   *deduplicator.Deduplicator[uint64]
	clock    *lamport.Clock
}

func New(r *replicas.Replicas, c *lamport.Clock, hostname string) *Counter {
	incr := &Counter{
		hostname: hostname,
		Replicas: r,
		dedupe:   deduplicator.New[uint64](1 * time.Minute),
		clock:    c,
	}

	r.Router.POST(Path, incr.handleIncrement)

	return incr
}

// Get returns the current counter value.
func (counter *Counter) Get() uint64 {
	return counter.counter.Load()
}

// Time returns the current Lamport clock value.
func (counter *Counter) Time() uint64 {
	return counter.clock.Get()
}

func (counter *Counter) Increment() {
	counter.process(counter.clock.Get())
}

// process ensures deduplication, broadcasting, and local increment of the counter.
func (incr *Counter) process(dedupeId uint64) bool {
	if incr.dedupe.Mark(dedupeId) {
		incr.clock.Set(dedupeId)

		if <-incr.broadcast() {
			incr.counter.Add(1)
			return true
		}
	}

	return false
}

// broadcast sends an increment request to all replicas, including this replica.
func (counter *Counter) broadcast() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := counter.Request().
			SetHeader("X-Deduplication-Id", counter.clock.String()).
			SetHeader("X-Replica-Id", counter.hostname)

		result := <-broadcast.
			Post(req, counter.Urls(Path)...).
			OnSuccess(func(resp *resty.Response) {
				fmt.Printf("Broadcast successful: %s\n", resp.Request.URL)
			}).
			OnError(func(err error) {
				fmt.Printf("Broadcast error: %v\n", err)
			}).
			SendC()

		done <- result != nil
	}()

	return done
}

// handleIncrement processes incoming increment requests to this replica
func (counter *Counter) handleIncrement(c *gin.Context) {
	dedupeID, err := strconv.ParseUint(c.GetHeader("X-Deduplication-Id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deduplication ID"})
		return
	}

	replicaID := c.GetHeader("X-Replica-Id")
	if replicaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing replica ID"})
		return
	}

	if !counter.process(dedupeID) {
		c.JSON(http.StatusOK, gin.H{"status": "already processed", "count": counter.Get()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "incremented", "count": counter.Get()})
}
