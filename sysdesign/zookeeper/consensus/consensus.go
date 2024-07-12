package consensus

import (
	"math/rand"
	"sgago/goginair-zookeeper/zknode"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	Path = "binary-problem"
)

var (
	value byte = 0
)

func Randomize() {
	value = byte(rand.Intn(2))
}

func Vote(conn *zk.Conn) error {
	acl := zk.WorldACL(zk.PermAll)

	_, _, err := zknode.CreateIfNotExists(conn, Path, []byte{}, zk.FlagEphemeral, acl)
	if err != nil {
		return err
	}

	_, err = conn.CreateProtectedEphemeralSequential(Path+"/participant-", []byte{value}, acl)
	if err != nil {
		return err
	}

	// event := <-watch

	return nil
}
