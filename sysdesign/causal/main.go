package main

import (
	"net/http"
	"sgago/thestudyguide-causal/config"
	"sgago/thestudyguide-causal/lamport"
	"sgago/thestudyguide-causal/replicas"
	"sgago/thestudyguide-causal/replicas/increment"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	homePath    = "/"
	healthzPath = "/healthz"
	incrPath    = "/increment"
)

func main() {
	router := gin.Default()
	resty := resty.New().SetTimeout(5 * time.Second)

	replicas := replicas.New(
		router,
		resty,
		config.MyId(),
		config.Hosts()...)

	router.GET(healthzPath, func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	router.GET(homePath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "hello",
		})
	})

	incr := increment.New(replicas, lamport.New())

	router.POST(incrPath, func(c *gin.Context) {
		<-incr.Do()
		c.Writer.WriteHeader(http.StatusOK)
	})

	go func() {
		for i := 0; i < 100; i++ {
			<-incr.Do()
			time.Sleep(10 * time.Second)
		}
	}()

	router.Run()
}
