package main

import (
	"net/http"
	"sgago/thestudyguide-causal/config"
	"sgago/thestudyguide-causal/replicas"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	home    = "/"
	healthz = "/healthz"
)

func main() {
	router := gin.Default()
	resty := resty.New()
	replicas := replicas.New(router, resty, config.OtherNames()...)

	router.GET(healthz, func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	router.GET(home, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "hello",
		})
	})

	go func() {
		for i := 0; i < 100; i++ {
			<-replicas.Healthz()
			time.Sleep(10 * time.Second)
		}
	}()

	router.Run()
}
