package counter

import (
	"log"
	"math/rand"
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
	syncPath = "/sync"

	IncrPath = replicas.GroupPath + incrPath
	SyncPath = replicas.GroupPath + syncPath
)

type Counter struct {
	*replicas.Replicas

	counter atomic.Uint64
	dedupe  *deduplicator.Deduplicator[uint64]
	clock   *lamport.Clock
}

func New(r *replicas.Replicas, c *lamport.Clock) *Counter {
	incr := &Counter{
		Replicas: r,
		dedupe:   deduplicator.New[uint64](1 * time.Minute),
		clock:    c,
	}

	r.Router.POST(IncrPath, incr.HandleIncrement)
	r.Router.GET(SyncPath, incr.HandleSync)

	go incr.sync()

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
	counter.process(counter.clock.Inc())
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
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Deduplication-Id", counter.clock.String()).
			SetHeader("X-Replica-Id", counter.Self())

		result := <-broadcast.
			Post(req, counter.Urls(IncrPath)...).
			Before(func(req *resty.Request) {
				log.Printf("Broadcasting to: %s\n", req.URL)
			}).
			OnSuccess(func(resp *resty.Response) {
				log.Printf("Broadcast successful: %s\n", resp.Request.URL)
			}).
			OnNetError(func(err error) {
				log.Printf("Broadcast network error: %v\n", err)
			}).
			OnAppError(func(resp *resty.Response) {
				log.Printf("Broadcast application error: %v\n", resp.Error())
			}).
			SendC()

		done <- result != nil
	}()

	return done
}

// HandleIncrement processes incoming increment requests to this replica
func (counter *Counter) HandleIncrement(c *gin.Context) {
	dedupeID, err := strconv.ParseUint(c.GetHeader("X-Deduplication-Id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deduplication ID"})
		return
	}

	replicaID := c.GetHeader("X-Replica-Id")
	if replicaID == "" {
		log.Println("Error: missing replica ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing replica ID"})
		return
	}

	if !counter.process(dedupeID) {
		log.Printf("Already processed dedupeId %d, count is %d, time is %d", dedupeID, counter.Get(), counter.Time())
		c.JSON(http.StatusOK, gin.H{"status": "already processed", "count": counter.Get()})
		return
	}

	log.Printf("Incremented! dedupeId %d, count is %d, time is %d", dedupeID, counter.Get(), counter.Time())
	c.JSON(http.StatusOK, gin.H{"status": "incremented", "count": counter.Get()})
}

// TODO: Finish anti-entropy mechanism. Should not be random. Could do all or round-robin.
// sync is an anti-entropy mechanism to synchronize the counter value with other replicas.
func (counter *Counter) sync() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		urls := counter.OtherUrls(SyncPath)

		if len(urls) == 0 {
			return
		}

		url := urls[rand.Intn(len(urls))]

		go func(u string) {
			log.Println("Random url:", u)

			syncResp := &syncResp{}
			resp, err := counter.RestyCli.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("X-Deduplication-Id", counter.clock.String()).
				SetHeader("X-Replica-Id", counter.Self()).
				SetResult(syncResp).
				Get(u)

			if err != nil {
				log.Printf("Sync error: %v\n", err)
				return
			}

			if resp != nil {
				log.Printf("Sync received %s, %+v\n", resp.Request.URL, syncResp)

				if syncResp.Time > counter.Time() {
					log.Printf("Out of sync, count is %d and time is %d\n", syncResp.Count, syncResp.Time)

					counter.counter.Store(syncResp.Count)
					counter.clock.Set(syncResp.Time)
				} else {
					log.Printf("In sync, count is %d and time is %d\n", counter.Get(), counter.Time())
				}
			}
		}(url)
	}
}

type syncResp struct {
	Count uint64 `json:"count"`
	Time  uint64 `json:"time"`
}

func (counter *Counter) HandleSync(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": counter.Get(),
		"time":  counter.Time(),
	})
}
