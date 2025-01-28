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
	Router *gin.Engine
	Client resty.Client
	consul *consul.Consul
	self   string
}

func New(router *gin.Engine, resty *resty.Client, consul *consul.Consul, self string) *Replicas {
	replicas := &Replicas{
		Router: router,
		Client: *resty,
		consul: consul,
		self:   self,
	}

	return replicas
}

func (r *Replicas) Request() *resty.Request {
	return r.Client.R()
}

func (r *Replicas) Services() []string {
	return r.consul.Services()
}

func (r *Replicas) Self() string {
	return r.self
}

func (r *Replicas) Others() []string {
	services := r.consul.Services()
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
