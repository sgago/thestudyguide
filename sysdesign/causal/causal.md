# Causal Consistency
Causal consistency is a weaker form of strong consistency where operations are applied in the same order by each of the replicas. An operation would be like "increment page count by one" or "the user pressed the delete key". Operations are sent to all nodes with a timestamp and the each node applies the operation in the correct order.
- Unlike strong consistency, the replicas do not wait for consensus which boosts performance. Nodes may disagree about values temporarily while updates are in-flight.
- Unlike eventual consistency, we maintain a global order of operations. We can avoid some of the confusing issues and bugs brought about by eventual consistency.
- There's less control plane issues to worry about or dancing around with raft/paxos consensus. However, since microservices are provisioned dynamically, they still need to find each other: hardcoded IP addresses, service registry like Consul, or DNS based like Traefik.
- Replicas will need to how to apply updates if they are missing or arrive out-of-order.
- Causal gives us availability and partition tolerance.

Causal is often considered good enough for many distributed applications. The CALM theorem and CRDTs give us helpful tools to reason about and reach causal consistency.

The basic idea with causal is to send updates to all nodes. Nodes will then apply the update in the correct order. Reads can be replied to immediately. Create, update, and delete operations require some tricks and tools.
- We will mirror create, update, and delete web requests to all replicas immediately. We will literally send one web request to each replica. This is a broadcast.
- Each replica will re-broadcast the same request to all other replicas. Having all the replica echo the request to all other replicas makes this broadcast more reliable and is called _eager reliable broadcast_.
- Each replica will re-send the request to itself. This can simplify the code some.
- Each replica will need to de-duplicate the message by a special deduplication Id. The reliable broadcast means each replica will receive the same message repeatedly. If you don't deduplicate the Ids, the replicas will never stop sending messages.
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