package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sgago/thestudyguide-raft/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

const (
	// URL Paths
	homePath    = "/"
	healthzPath = "/healthz"
	leaderPath  = "/leader"
	statePath   = "/state"
	addNodePath = "/add_node"
)

var raftNode *raft.Raft
var valueFsm *valueFSM

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %v", err)
	}

	raftAddr := config.RaftAddress()

	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(raftAddr)

	store, err := raftboltdb.NewBoltStore(hostname + "-raft-store.db")
	if err != nil {
		log.Fatalf("Failed to create BoltDB store: %v", err)
	}

	logStore, err := raft.NewLogCache(512, store)
	if err != nil {
		log.Fatalf("Failed to create log store: %v", err)
	}

	snapshotStore, err := createSnapshotStoreWithRetry(".", 3, 500*time.Millisecond)
	if err != nil {
		log.Fatalf("Failed to create snapshot store: %v", err)
	}

	transport, err := raft.NewTCPTransport(raftAddr, nil, 2, 10*time.Second, os.Stderr)
	if err != nil {
		log.Fatalf("Failed to create TCP transport: %v", err)
	}

	valueFsm = &valueFSM{}

	raftNode, err = raft.NewRaft(c, valueFsm, logStore, store, snapshotStore, transport)
	if err != nil {
		log.Fatalf("Failed to create Raft node: %v", err)
	}

	nodes := config.RaftNodes()
	var servers []raft.Server

	for _, node := range nodes {
		servers = append(servers, raft.Server{
			Suffrage: raft.Voter,
			ID:       raft.ServerID(node),
			Address:  raft.ServerAddress(node),
		})
	}

	cfg := raft.Configuration{
		Servers: servers,
	}

	f := raftNode.BootstrapCluster(cfg)
	if err := f.Error(); err != nil {
		if !strings.Contains(err.Error(), "bootstrap only works on new clusters") {
			fmt.Print(fmt.Errorf("raft.Raft.BootstrapCluster: %v", err))
		}
	}

	r := gin.Default()

	r.GET(healthzPath, healthz)
	r.POST(addNodePath, addNode)

	r.POST("/set_value", func(c *gin.Context) {
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		leader := raftNode.Leader()
		if leader == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No leader available"})
			return
		}

		if leader != transport.LocalAddr() {
			host, _, _ := net.SplitHostPort(string(leader))

			// Forward the request to the leader node
			url := fmt.Sprintf("http://%s:%s/set_value", host, "8080")

			fmt.Println("Forwarding request to leader:", url)

			cli := resty.New()

			resp, err := cli.R().
				EnableTrace().
				SetBody(map[string]string{"Value": json.Value}).
				Post(url)

			fmt.Println(resp.Request.GenerateCurlCommand())

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(resp.StatusCode(), gin.H{"message": resp.String()})

			return
		}

		if err := setValue(json.Value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "value set"})
	})

	r.GET("/get_value", func(c *gin.Context) {
		if valueFsm == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "FSM not initialized"})
			return
		}
		value := valueFsm.GetValue()
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.GET(homePath, func(c *gin.Context) {
		leader := raftNode.Leader()
		state := raftNode.State().String()

		c.JSON(http.StatusOK, gin.H{
			"hostname": hostname,
			"raftAddr": raftAddr,
			"leader":   leader,
			"state":    state,
		})
	})

	r.Run()
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func createSnapshotStoreWithRetry(dir string, retries int, delay time.Duration) (raft.SnapshotStore, error) {
	var snapshotStore raft.SnapshotStore
	var err error
	for i := 0; i < retries; i++ {
		snapshotStore, err = raft.NewFileSnapshotStore(dir, 1, os.Stderr)
		if err == nil {
			return snapshotStore, nil
		}
		log.Printf("Failed to create snapshot store (attempt %d/%d): %v", i+1, retries, err)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("failed to create snapshot store after %d attempts: %v", retries, err)
}

type SetValueCommand struct {
	Value string `json:"value" binding:"required"`
}

func setValue(value string) error {
	command := SetValueCommand{Value: value}
	data, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("failed to marshal command: %v", err)
	}

	// Apply the command to the Raft log
	f := raftNode.Apply(data, 10*time.Second)
	if err := f.Error(); err != nil {
		return fmt.Errorf("failed to apply command: %v", err)
	}

	return nil
}

func addNode(c *gin.Context) {
	var json struct {
		NodeID  string `json:"node_id" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idxFuture := raftNode.AddVoter(raft.ServerID(json.NodeID), raft.ServerAddress(json.Address), 0, 0)
	if idxFuture.Error() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": idxFuture.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "node added"})
}
