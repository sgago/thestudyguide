package counter

import (
	"log"
	"math/rand"
	"net/http"
	"sgago/thestudyguide-causal/broadcast"
	"sgago/thestudyguide-causal/dedupe"
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
	dedupe  *dedupe.Dedupe[uint64]
	clock   *lamport.Clock
}

func New(r *replicas.Replicas, c *lamport.Clock) *Counter {
	incr := &Counter{
		Replicas: r,
		dedupe:   dedupe.New[uint64](1 * time.Minute),
		clock:    c,
	}

	r.Router.POST(IncrPath, incr.HandleIncrement)
	r.Router.GET(SyncPath, incr.HandleSync)

	go incr.StartSync()

	return incr
}

// Get returns the current counter value.
func (c *Counter) Get() uint64 {
	return c.counter.Load()
}

// Time returns the current Lamport clock value.
func (c *Counter) Time() uint64 {
	return c.clock.Get()
}

func (c *Counter) Increment() {
	c.process(c.clock.Inc())
}

// process ensures deduplication, broadcasting, and local increment of the counter.
func (c *Counter) process(dedupeId uint64) bool {
	if c.dedupe.Mark(dedupeId) {
		c.clock.Set(dedupeId)

		if <-c.broadcast() {
			c.counter.Add(1)
			return true
		}
	}

	return false
}

// broadcast sends an increment request to all replicas, including this replica.
func (c *Counter) broadcast() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := c.Request().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Deduplication-Id", c.clock.String()).
			SetHeader("X-Replica-Id", c.Self())

		result := <-broadcast.
			Post(req, c.Urls(IncrPath)...).
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
func (c *Counter) HandleIncrement(ctx *gin.Context) {
	dedupeID, err := strconv.ParseUint(ctx.GetHeader("X-Deduplication-Id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid deduplication ID"})
		return
	}

	replicaID := ctx.GetHeader("X-Replica-Id")
	if replicaID == "" {
		log.Println("Error: missing replica ID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing replica ID"})
		return
	}

	if !c.process(dedupeID) {
		log.Printf("Already processed dedupeId %d, count is %d, time is %d", dedupeID, c.Get(), c.Time())
		ctx.JSON(http.StatusOK, gin.H{"status": "already processed", "count": c.Get()})
		return
	}

	log.Printf("Incremented! dedupeId %d, count is %d, time is %d", dedupeID, c.Get(), c.Time())
	ctx.JSON(http.StatusOK, gin.H{"status": "incremented", "count": c.Get()})
}

// StartSync is an anti-entropy mechanism to synchronize the counter value with other replicas.
func (c *Counter) StartSync() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		urls := c.OtherUrls(SyncPath)

		if len(urls) == 0 {
			return
		}

		url := urls[rand.Intn(len(urls))]

		go func(u string) {
			log.Println("Random url:", u)

			syncResp := &syncResp{}
			resp, err := c.RestyCli.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("X-Deduplication-Id", c.clock.String()).
				SetHeader("X-Replica-Id", c.Self()).
				SetResult(syncResp).
				Get(u)

			if err != nil {
				log.Printf("Sync error: %v\n", err)
				return
			}

			if resp != nil {
				log.Printf("Sync received %s, %+v\n", resp.Request.URL, syncResp)

				if syncResp.Time > c.Time() {
					log.Printf("Out of sync, count is %d and time is %d\n", syncResp.Count, syncResp.Time)

					c.counter.Store(syncResp.Count)
					c.clock.Set(syncResp.Time)
				} else {
					log.Printf("In sync, count is %d and time is %d\n", c.Get(), c.Time())
				}
			}
		}(url)
	}
}

type syncResp struct {
	Count uint64 `json:"count"`
	Time  uint64 `json:"time"`
}

func (c *Counter) HandleSync(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"count": c.Get(),
		"time":  c.Time(),
	})
}
