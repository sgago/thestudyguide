package main

import (
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
	var consulCli consul.Client
	var err error

	if !envcfg.IsLocal() {
		consulCli, err = consul.New()
		if err != nil {
			log.Fatalf("Failed to create Consul client: %v", err)
		}

		if err := consulCli.Register(); err != nil {
			log.Fatalf("Failed to register service with Consul: %v", err)
		}
	} else {
		consulCli = consul.NewMock()
	}

	consulCli.StartTTL()

	router := gin.Default()

	router.GET(healthPath, func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	router.GET(homePath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "hello",
		})
	})

	restyCli := resty.New().SetTimeout(15 * time.Second)

	// TODO: Chaining like this ugly, make a better counter Ctor
	replicas := replicas.New(router, restyCli, consulCli)
	counter := counter.New(replicas, lamport.New())

	router.GET(incrPath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"replica": envcfg.HostName(),
			"count":   counter.Get(),
			"time":    counter.Time(),
		})
	})

	router.POST(incrPath, func(c *gin.Context) {
		counter.Increment()

		c.Writer.WriteHeader(http.StatusOK)

		c.JSON(http.StatusOK, gin.H{
			"replica": envcfg.HostName(),
			"count":   counter.Get(),
			"time":    counter.Time(),
		})
	})

	go func() {
		log.Println("Starting server for", envcfg.HostName())

		if err := router.Run(); err != nil {
			log.Fatalf(("Failed to start server: %v"), err)
		}
	}()

	select {}
}
