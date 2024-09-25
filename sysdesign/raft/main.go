package main

import (
	"encoding/json"
	"errors"
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

var raftNode *raft.Raft

var hostname string
var addr string
var bolt *raftboltdb.BoltStore
var logs *raft.LogCache
var snaps raft.SnapshotStore
var trans *raft.NetworkTransport

var fsm *valueFSM

func main() {
	var err error

	hostname, err = os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %v", err)
	}

	addr = config.RaftAddress()

	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(addr)

	bolt, err = raftboltdb.NewBoltStore(hostname + "-raft-store.db")
	if err != nil {
		log.Fatalf("Failed to create BoltDB store: %v", err)
	}

	// logCache is a replicated log that tracks cluster changes.
	// It is used to ensure that the cluster configuration is consistent across all nodes.
	// For example, when new data are added, the log store adds a corresponding log entry.
	// The log store stores opaque binary blobs, so it is up to the application to interpret the data.
	// This FSM uses Gob encoding to serialize and deserialize the data.
	logs, err = raft.NewLogCache(512, bolt)
	if err != nil {
		log.Fatalf("Failed to create log store: %v", err)
	}

	// snaps is used to store snapshots of the FSM.
	// It minimizes disk usage and helps us avoid replaying the entire log over and over again.
	// The logStore can grow forever, so the snapshotStore is used to store snapshots of the FSM
	// and let's us compact (remove) old log entries to save space.
	snaps, err = createSnapshotStoreWithRetry(".", 3, 500*time.Millisecond)
	if err != nil {
		log.Fatalf("Failed to create snapshot store: %v", err)
	}

	// trans (transport) is used to setup communicate with other nodes in the cluster.
	// Here we use good old TCP though gRPC is also well supported.
	trans, err = raft.NewTCPTransport(addr, nil, 2, 10*time.Second, os.Stderr)
	if err != nil {
		log.Fatalf("Failed to create TCP transport: %v", err)
	}

	// valueFsm is the finite state machine that stores the value, applies log entries, and creates snapshots.
	// This Raft FSM stores a single JSON string value called "value" like {"value": "foo"}.
	// The JSON will be Gob encoded and decoded.
	fsm = &valueFSM{}

	// Create the raft node with the configuration, log store, and FSM.
	// The node is the main entry point for the Raft protocol.
	// It is responsible for coordinating the consensus protocol and applying log entries to the FSM.
	// If you need to add more state machines, you can create a new FSM and call NewRaft again.
	raftNode, err = raft.NewRaft(c, fsm, logs, bolt, snaps, trans)
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

	// Next, we need to bootstrap the cluster. This is tricky because we only want one node to do this once.
	// However, we definitely need one node to bootstrap the cluster or we won't have a functioning raft cluster!
	// We could use env var tricks or a distributed lock to ensure only one node bootstraps the cluster.
	// Instead, we will just have every node try to bootstrap the cluster and safely ignore the "cluster already exists"
	// error when it fails. BootstrapCluster returns a future; we wait on that future via the Error method.
	future := raftNode.BootstrapCluster(cfg)
	if err := future.Error(); err != nil {
		if !strings.Contains(err.Error(), "bootstrap only works on new clusters") {
			fmt.Print(fmt.Errorf("raft.Raft.BootstrapCluster: %v", err))
		}
	}

	r := gin.Default()

	r.GET("/", stateHandler)
	r.GET("/healthz", healthz)
	r.POST("/node", addNodeHandler)
	r.POST("/value", setValueHandler)
	r.GET("/value", getValueHandler)

	r.Run()
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func stateHandler(c *gin.Context) {
	leader := raftNode.Leader()
	state := raftNode.State().String()

	c.JSON(http.StatusOK, gin.H{
		"hostname": hostname,
		"raftAddr": addr,
		"leader":   leader,
		"state":    state,
	})
}

func getValueHandler(c *gin.Context) {
	if fsm == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FSM not initialized"})
		return
	}
	value := fsm.GetValue()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

func leader() (raft.ServerAddress, bool, error) {
	leader := raftNode.Leader()
	if leader == "" {
		return "", false, errors.New("no leader is available")
	}

	return leader, leader == trans.LocalAddr(), nil
}

func setValueHandler(c *gin.Context) {
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leader, isLeader, err := leader()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No leader available"})
		return
	}

	if !isLeader {
		// Only the leader can accept writes and new nodes to the cluster.
		// Therefore, we will forward the request to the leader node.
		// Write attempts by a non-leader will cause an error.

		host, _, _ := net.SplitHostPort(string(leader))
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

func addNodeHandler(c *gin.Context) {
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
