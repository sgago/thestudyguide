package main

import (
	"net/http"
	"sgago/thestudyguide-causal/config"
	"sgago/thestudyguide-causal/lamport"
	"sgago/thestudyguide-causal/replicas"

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
	resty := resty.New()
	clock := lamport.New()

	replicas := replicas.New(
		router,
		resty,
		clock,
		config.OtherNames()...)

	router.GET(healthzPath, func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	router.GET(homePath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "hello",
		})
	})

	router.POST(incrPath, func(c *gin.Context) {
		<-replicas.BroadcastIncrement()
		c.Writer.WriteHeader(http.StatusOK)
	})

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		<-replicas.Healthz()
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	router.Run()
}
