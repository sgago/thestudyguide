# Causal Consistency
Causal consistency is a weaker form of strong consistency where the order of operations are applied in the same order. An operation would be like "increment page count" or "user A pressed the delete key". These operations are sent with a timestamp and the nodes apply the operations in the correct order. It's often considered good enough for many distributed applications.
- Unlike strong consistency, the replicas do not wait for consensus which boosts performance. Nodes may disagree about values temporarily while updates are in-flight.
- Unlike eventual consistency, we maintain a global order of operations. We can avoid some of the confusing issues and bugs brought about by eventual consistency.
- There's no control plane to worry about or dancing with raft/paxos implementations.
- Replicas will need to how to apply updates if they are missing or arrive out-of-order.
- Causal gives us availability and partition tolerance.

The basic idea with causal is to send updates to all nodes. Nodes will then apply the update in the correct order. Reads can be replied to immediately. Create, update, and delete operations require some tricks and tools.
- We will mirror create, update, and delete web requests to all replicas immediately. We will literally send one web request to each replica. This is a broadcast.
- Each replica will re-broadcast the same request to all other replicas. This is known as an eager reliable broadcast. Having all the replica echo the request to all other replicas makes this "reliable".
- Each replica will re-send the request to itself. This maybe simplifies the code a bit.
- Each replica will need to de-duplicate the message by a special deduplication Id. The reliable broadcast means each replica will receive the same message over and over again.
- We will use a logical clock, a lamport or vector clock, for the deduplication Id sent with each message.
- Each node will use the lamport clock value to order events appropriately, if they even need to.

So, again, broadcast requests to all replicas, deduplicate the requests by a lamport clock value, and apply the operations in order of the clock.

## Reliable broadcast

### Eager broadcast

### Gossip broadcast

### Other broadcasts

## Logical clocks

### Lamport clocks

### Vector clocks

### Logical clocks for deduplicating messages

## The CALM Theorem

## Conflict-free replicated data types

### Increment

### Last writer wins

### Multi-value

### Replicated log model

## Repair mechanisms

## Compaction or Deletion