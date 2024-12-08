package healthz

import (
	"fmt"
	"net/http"
	"sgago/thestudyguide-causal/broadcast"
	"sgago/thestudyguide-causal/replicas"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	healthzPath = "/healthz"

	HealthzPath = replicas.GroupPath + healthzPath
)

type Healthz struct {
	*replicas.Replicas
}

func New(r *replicas.Replicas) *Healthz {
	h := &Healthz{
		Replicas: r,
	}

	r.Router.GET(HealthzPath, h.handle)

	return h
}

func (r *Healthz) Do() <-chan bool {
	return r.broadcast()
}

func (r *Healthz) broadcast() <-chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		req := r.Request()
		urls := r.Urls(HealthzPath)

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

func (r *Healthz) handle(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
