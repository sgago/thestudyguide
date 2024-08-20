# Apache Zookeeper

Zookeeper (ZK) is a distributed coordination service. It provides:
- Configuration management
- Leader election
- Distributed locks
- Cluster management

ZK usage is based around Znodes that are stored in the control plane. The Znodes (ZN) act a lot like folders in a file system. You can create, update, and delete the Znodes. You can also store data within the Znode themselves. ZN have a hierarchy like folders and they even use `/` for pathing:
- `/myZnode`
- `/myZnode/myChildZnode`
- `/election/leader-00000001`
- `/election/leader-00000002`

For example, we can store JSON data in the `/configuration` Znode that workers read on startup. You can also set the ACL on Znodes so that they aren't accidentally deleted or updated.

Our worker nodes can also *watch* Znodes for changes. For example, maybe workers watch the configuration Znode for changes and immediately update configuration values.

Obviously, there are some differences between folders and Znodes. Znodes are used for distributed coordination, so there are different types of them.
- Persistent - A persistent ZN stays around "forever", or until they are manually deleted. They are useful for cluster configuration.
- Ephemeral - Ephemeral or temporary Znodes do not last forever. If a client session ends, then that Znode is delete, too. Ping requests are used to determine if the session is still alive. We can use these for distributed locking where a single-node needs to perform work without being interrupted.
- Epehmeral Sequential - Simiilar to ephemeral but assigns sequentially increasing numbers to the Znode names. This can be used for leader election, consensus, and so on. The paths look like  `/myZnode-<sequential number>`, `/myZnode-00000001`, or `/election/leader-00000001`.
- Persistent sequential - Similar to ephemeral sequential but ZN stick around forever. Usage is rare.

## Zookeeper control plane
ZK requires its own control plane, usually we specifiy an odd number of nodes like 3, 5 or 7 for quorum (>50% of the vote). More control nodes should increase availability at the cost of money, coordination, etc. The control plane node has its own leader election and voting processes wholly separate from our worker nodes.

### Zookeeper control plane env vars and ports
Of note, you'll want to set each control plane nodes' ZOO_MY_ID and ZOO_SERVERS environment variables.

```yaml
environment:
    ZOO_MY_ID: 1
    ZOO_SERVERS: server.1=zoo1:2888:3888;2181 server.2=zoo2:2888:3888;2181 server.3=zoo3:2888:3888;2181
```

The important port numbers to be aware of and expose are:
```yaml
expose:
    - 2181 # Port 2181 is the client port, use this one to communicate
    - 2888 # Port 2888 is the quorum port, used by leader to communicate with followers (for control plane only)
    - 3888 # Port 3888 is the election port to vote for a leader (for control plane only)
```

## Zookeeper SDK
Fortunately, there are typically ZK SDKs for us to use. For example, Go devs can use [go-zookeeper](https://github.com/go-zookeeper/zk). It simplifies using ZK quite a bit.

## Zookeeper recipes
TODO: Author some ZK recipes so readers get a feel for how to use it. These help a lot.

## References
- A greater starter for learning ZK from Medium [here](https://bikas-katwal.medium.com/zookeeper-introduction-designing-a-distributed-system-using-zookeeper-and-java-7f1b108e236e)