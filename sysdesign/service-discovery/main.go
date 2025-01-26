package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "hello",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		// Respond with 200 OK and the message "OK"
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	go router.Run()
	go consul()

	select {}
}

func consul() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"

	// Step 1: Create a Consul client
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating Consul client: %v", err)
	}

	// Define a unique check ID
	checkID := "goginair-" + os.Getenv("HOSTNAME") + "-check" // Unique check ID

	// Step 2: Define the service registration
	registration := &consulapi.AgentServiceRegistration{
		ID:      fmt.Sprintf("goginair-%s", hostname), // Unique service ID
		Name:    "goginair",                           // Service name
		Port:    8080,                                 // Service port
		Address: "goginair",                           // Service address
		Tags:    []string{"http"},                     // Tags for metadata
		Check: &consulapi.AgentServiceCheck{ // Health check configuration
			CheckID:                        checkID,
			TTL:                            "10s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	// Step 3: Register the service with Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Error registering service with Consul: %v", err)
	}

	log.Println("Service registered successfully with Consul")

	// Update TTL periodically (simulate the health check)
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		log.Println("Updating TTL")

		// Simulate healthy TTL update
		err := client.Agent().UpdateTTL(checkID, "TTL OK", consulapi.HealthPassing)
		if err != nil {
			log.Printf("Error updating TTL: %v", err)
		}
	}
}
