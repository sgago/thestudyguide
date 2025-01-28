package health

import (
	"log"
	"net/http"
	"sgago/thestudyguide-causal/broadcast"
	"sgago/thestudyguide-causal/replicas"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	healthPath = "/health"

	HealthPath = replicas.GroupPath + healthPath
)

type Health struct {
	*replicas.Replicas
}

func New(r *replicas.Replicas) *Health {
	h := &Health{
		Replicas: r,
	}

	r.Router.GET(HealthPath, h.handle)

	return h
}

func (h *Health) Do() <-chan bool {
	return h.broadcast()
}

func (h *Health) broadcast() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := h.Request()
		urls := h.Urls(HealthPath)

		result := <-broadcast.
			Post(req, urls...).
			OnSuccess(func(resp *resty.Response) {
				log.Println(resp.Request.GenerateCurlCommand())
			}).
			OnNetError(func(err error) {
				log.Println(err)
			}).
			OnAppError(func(resp *resty.Response) {
				log.Println(resp.Error())
			}).
			SendC()

		done <- result != nil
	}()

	return done
}

func (h *Health) handle(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
