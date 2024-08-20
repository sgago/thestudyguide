# Messaging
Messaging is an indirect form of communication where messages, sometimes called events or records, are passed through a message broker. This abstraction brings some niceties at the cost of increased complexity and cost.

Brokers decouple senders and receivers allowing them to scale separately. The same message can be broadcast to different systems, e.g., sending the same message to a payment and shipping processors. The broker can also hold messages while the consumers is down without losing data.

Message brokers serve a variety of use cases: microservice communication, financial transactions, IoT sensor data, outbox patterns, saga patterns, and so on. There's lots of off-the-shelf message brokers: RabbitMQ, Azure Service Bus, and GCP Pub/Sub. This guide will examine Kafka.

## Kafka
[Kafka](https://kafka.apache.org/documentation/#gettingStarted) is a message broker. In Kafka parlance, its an event streaming platform. Events (messages) are published (written) and subscribed to (read).

### Events
Events are a record of something happening: a transaction, an order being placed, or even errors. Events contain both data and metadata, sort of like a file or an email. They have:
- Unique key or Id
- Data or value: often JSON or similar
- Timestamp

### Producers and Consumers
In short, producers create messages and consumers read those messages.
- A producer could be a checkout service for a digital store.
- A consumer could be the payment processor and inventory manager.
Again, producers do not need to wait for consumers before publishing more messages. Kafka also takes great pains to try and hit exactly-once processing in its closed system.

### Topics
If events are files then topics would be considered folders. Topics can have many producers (writters) and consumers (readers). Unlike traditional message brokers, messages can be read many times instead of being automatically deleted after being read. The topic configuration defines when events are deleted.

### Partitions
If topics are folders, then partitions are special subfolders or buckets. Partitions divide the messages to achieve scalability. For example, we might need three consumers to keep up with the all the messages from only one producer. In this case, we can create three partitions, one for each consumer, the Kafka will spread the messages to each partition in a round robin fashion.

### Batching
Producers and consumers can push and pull one message at a time; however, we often push and pull messages in batches. For example, a consumer can grab a batch of ten messages at a time. This increases the latency of individual messages but typically improves network utilization, compression ratios, and throughput and is usually a worthwhile tradeoff. Kafka can even be configured to wait to accumulate messages, buffer up some work for until a set amount of time is reached, before processing.

## References
- [Kafka Quickstart](https://kafka.apache.org/quickstart)
- [Kafka Design](https://kafka.apache.org/documentation/#majordesignelements)
- [Kafdrop](https://github.com/obsidiandynamics/kafdrop)