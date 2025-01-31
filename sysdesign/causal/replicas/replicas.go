package replicas

import (
	"net/url"
	"sgago/thestudyguide-causal/consul"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	GroupPath = "/causal"
)

type Replicas struct {
	Router    *gin.Engine
	RestyCli  resty.Client
	ConsulCli consul.Client
}

func New(router *gin.Engine, restyCli *resty.Client, consulCli consul.Client) *Replicas {
	replicas := &Replicas{
		Router:    router,
		RestyCli:  *restyCli,
		ConsulCli: consulCli,
	}

	return replicas
}

func (r *Replicas) Request() *resty.Request {
	return r.RestyCli.R()
}

func (r *Replicas) Services() []string {
	return r.ConsulCli.Services()
}

func (r *Replicas) Self() string {
	return r.ConsulCli.Self()
}

func (r *Replicas) Others() []string {
	services := r.ConsulCli.Services()
	self := r.Self()
	var others []string

	for _, service := range services {
		if service != self {
			others = append(others, service)
		}
	}

	return others
}

func (r *Replicas) Urls(path ...string) []string {
	var urls []string

	for _, host := range r.Services() {
		u, _ := url.JoinPath(host, path...)
		urls = append(urls, u)
	}

	return urls
}

func (r *Replicas) OtherUrls(path ...string) []string {
	var urls []string

	for _, host := range r.Others() {
		u, _ := url.JoinPath(host, path...)
		urls = append(urls, u)
	}

	return urls
}
