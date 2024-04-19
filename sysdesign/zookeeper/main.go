package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	pingPath    = "/ping"
	healthzPath = "/healthz"
	nodePath    = "/node"

	zkNode = "localhost:21811"
	zkPath = "/test"
)

func main() {
	r := gin.Default()

	r.GET(healthzPath, healthz)
	r.GET(pingPath, ping)
	r.GET(nodePath, node)

	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func node(c *gin.Context) {
	conn, _, err := zk.Connect([]string{"zoo1:2181"}, time.Second*10)
	if err != nil {
		panic(err)
	}

	children, stat, _ /*ch*/, err := conn.ChildrenW("/")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v %+v\n", children, stat)
}
