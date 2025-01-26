package replicas

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	GroupPath = "/causal"
)

type Replicas struct {
	Router *gin.Engine
	client resty.Client
	hosts  []string
}

func New(router *gin.Engine, resty *resty.Client, hosts ...string) *Replicas {
	replicas := &Replicas{
		Router: router,
		client: *resty,
		hosts:  hosts,
	}

	return replicas
}

func (r *Replicas) Request() *resty.Request {
	return r.client.R()
}

func (r *Replicas) Urls(path ...string) []string {
	var urls []string

	for _, host := range r.hosts {
		u, _ := url.JoinPath(host, path...)
		urls = append(urls, u)
	}

	return urls
}
