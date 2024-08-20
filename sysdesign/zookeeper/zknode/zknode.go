package zknode

import "github.com/samuel/go-zookeeper/zk"

// CreateIfNotExists creates a Znode if it does not exist.
func CreateIfNotExists(conn *zk.Conn, path string, data []byte, flags int32, acl []zk.ACL) (bool, *zk.Stat, error) {
	exists, stat, err := conn.Exists(path)
	if err != nil {
		return false, nil, err
	} else if exists {
		return false, stat, nil
	}

	_, err = conn.Create(path, data, flags, acl)
	if err != nil {
		if err != zk.ErrNodeExists {
			return false, nil, err
		}

		_, stat, err := conn.Exists(path)
		if err != nil {
			return false, nil, err
		}

		return false, stat, nil
	}

	return true, stat, nil
}

// CreateIfNotExists creates a Znode if it does not exist and returns
// a watch chan to monitor any changes.
func CreateIfNotExistsW(conn *zk.Conn, path string, data []byte, flags int32, acl []zk.ACL) (bool, *zk.Stat, <-chan zk.Event, error) {
	exists, stat, watch, err := conn.ExistsW(path)
	if err != nil {
		return false, nil, nil, err
	} else if exists {
		return false, stat, watch, nil
	}

	_, err = conn.Create(path, data, flags, acl)
	if err != nil {
		if err != zk.ErrNodeExists {
			return false, nil, nil, err
		}

		_, stat, watch, err := conn.ExistsW(path)
		if err != nil {
			return false, nil, nil, err
		}

		return false, stat, watch, nil
	}

	return true, stat, watch, nil
}
