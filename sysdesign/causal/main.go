package main

import (
	"fmt"
	"log"
	"net/http"
	"sgago/thestudyguide-causal/consul"
	"sgago/thestudyguide-causal/envcfg"
	"sgago/thestudyguide-causal/lamport"
	"sgago/thestudyguide-causal/replicas"
	"sgago/thestudyguide-causal/replicas/counter"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	homePath   = "/"
	healthPath = "/health"
	incrPath   = "/increment"
)

func main() {
	consul, err := consul.New()
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	if err := consul.Register(); err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}

	consul.StartTTL()

	router := gin.Default()

	router.GET(healthPath, func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	router.GET(homePath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "hello",
		})
	})

	resty := resty.New().SetTimeout(15 * time.Second)

	// TODO: This does not get a current list of replicas from Consul.
	replicas := replicas.New(
		router,
		resty,
		consul.Services()...)

	counter := counter.New(replicas, lamport.New(), envcfg.HostName())

	router.POST(incrPath, func(c *gin.Context) {
		fmt.Println(("received increment request"))
		counter.Increment()
		c.Writer.WriteHeader(http.StatusOK)

		//TODO: Put a body here so we can see the count
	})

	go func() {
		fmt.Println("Starting server")
		if err := router.Run(); err != nil {
			log.Fatalf(("Failed to start server: %v"), err)
		}
	}()

	select {}
}
