package leader

import (
	"fmt"
	"log"
	"sgago/goginair-zookeeper/participant"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	Path = "/leader"
)

var (
	isLeader = false
)

// IsLeader returns true if this Znode is the leader; otherwise, false.
func IsLeader() bool {
	return isLeader
}

func Election(conn *zk.Conn) {
	for {
		acl := zk.WorldACL(zk.PermRead | zk.PermDelete)

		if _, err := conn.Create(Path, participant.Bson(), zk.FlagEphemeral, acl); err == nil {
			isLeader = true // This node was able to create the znode, so it's the leader
		} else {
			isLeader = false // This node failed to create the znode, it's not a leader
		}

		fmt.Println("isLeader:", isLeader)

		if exists, _, watch, err := conn.ExistsW(Path); err != nil {
			log.Fatal(err)
		} else if exists {
			<-watch
		}
	}
}
