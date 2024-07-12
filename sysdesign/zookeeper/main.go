package main

import (
	"fmt"
	"log"
	"net/http"
	"sgago/goginair-zookeeper/config"
	"sgago/goginair-zookeeper/leader"
	"sgago/goginair-zookeeper/participant"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	// URL Paths
	pingPath         = "/ping"
	healthzPath      = "/healthz"
	nodePath         = "/node"
	leaderPath       = "/leader"
	participantsPath = "/participants"
)

func main() {
	r := gin.Default()

	r.GET(healthzPath, healthz)
	r.GET(nodePath, node)
	r.GET(leaderPath, isLeader)
	r.GET(participantsPath, allParticipants)

	conn, _ /*event*/, err := zk.Connect(config.ZkServers(), time.Second*10)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	if err := participant.Join(conn); err != nil {
		log.Fatal(err)
	}

	go leader.Election(conn)

	r.Run()
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func isLeader(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hostName": config.HostName(),
		"isleader": leader.IsLeader(),
	})
}

func allParticipants(c *gin.Context) {
	conn, _ /*event*/, err := zk.Connect(config.ZkServers(), time.Second*10)
	if err != nil {
		panic(err)
	}

	participants, err := participant.All(conn)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"participants": participants,
	})
}

func node(c *gin.Context) {
	conn, _ /*event*/, err := zk.Connect(config.ZkServers(), time.Second*10)
	if err != nil {
		panic(err)
	}

	children, stat, _ /*ch*/, err := conn.ChildrenW("/")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v %+v\n", children, stat)

	c.JSON(http.StatusOK, gin.H{
		"children": children,
		"stat":     stat,
	})
}
