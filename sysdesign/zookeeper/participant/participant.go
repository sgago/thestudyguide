package participant

import (
	"fmt"
	"sgago/goginair-zookeeper/config"

	"github.com/samuel/go-zookeeper/zk"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	Path = "/participants"
)

type Participant struct {
	HostName string `bson:"host"`
}

var (
	self []byte
)

func init() {
	data, err := bson.Marshal(Participant{
		HostName: config.HostName(),
	})

	if err != nil {
		panic(err)
	}

	self = data
}

func Bson() []byte {
	return self
}

func setup(conn *zk.Conn) error {
	acl := zk.WorldACL(zk.PermAll)

	if exists, _, err := conn.Exists(Path); err != nil {
		return err
	} else if !exists {
		if _, err := conn.Create(Path, []byte{}, 0, acl); err != nil {
			return err
		}
	}

	return nil
}

func Join(conn *zk.Conn) error {
	if err := setup(conn); err != nil {
		return err
	}

	acl := zk.WorldACL(zk.PermAll)

	path := Path + "/participant-"

	if _, err := conn.CreateProtectedEphemeralSequential(path, self, acl); err != nil {
		return err
	}

	fmt.Println("Joined with znode", path)

	return nil
}

func All(conn *zk.Conn) ([]Participant, error) {
	participants := make([]Participant, 0, 3)

	children, _, err := conn.Children(Path)

	if err != nil {
		return nil, err
	}

	fmt.Println("children paths:", children)

	for _, child := range children {
		path := Path + "/" + child
		fmt.Println("child path", path)

		data, _, err := conn.Get(path)

		if err != nil {
			return nil, err
		}

		var participant Participant
		err = bson.Unmarshal(data, &participant)
		if err != nil {
			return nil, err
		}

		fmt.Println("participant data", participant)

		participants = append(participants, participant)
	}

	return participants, nil
}
