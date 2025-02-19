# Messaging
Messaging is an indirect form of communication where messages, sometimes called events or records, are passed through a message broker. This abstraction brings some niceties at the cost of increased complexity and cost.

Brokers decouple senders and receivers allowing them to scale separately. The same message can be broadcast to different systems, e.g., sending the same message to a payment and shipping processors. The broker can also hold messages while consumers are down without losing data.

Message brokers serve a variety of use cases: microservice communication, financial transactions, IoT sensor data, outbox patterns, saga patterns, and so on. There's lots of off-the-shelf message brokers: RabbitMQ, Azure Service Bus, and GCP Pub/Sub. This guide will examine Kafka.

## Kafka
[Kafka](https://kafka.apache.org/documentation/#gettingStarted) is a message broker. In Kafka parlance, it's an event streaming platform. Events (messages) are published (written) and subscribed to (read).

### Events
Events are a record of something happening: a transaction, an order being placed, or even errors. Events contain both data and metadata, sort of like files, emails, and web requests have both a body and metadata. Events in Kafka have:
- A unique key or Id
- Some data: often JSON or similar
- A timestamp

### Producers and Consumers
In short, producers create messages and consumers read those messages.
- A producer could be a checkout service for a digital store. When a customer buys something, the checkout service writes an event. 
- A consumer could be the payment processor and inventory manager. When a purchase event is read, it is then processed by the payment processor.
Again, producers do not need to wait for consumers before publishing more messages. Kafka also takes great pains to try and hit exactly-once processing (within its closed system).

### Topics
If events are files then topics are like folders. Topics can have many producers (writers) and consumers (readers). Unlike traditional message brokers, messages can be read many times instead of being automatically deleted after being read. The topic configuration defines when events are deleted.

### Partitions
If topics are folders, then partitions are like sub-folders or buckets. Partitions divide the messages to achieve scalability. For example, we might need three consumers to keep up with the all the messages from only one producer. In this case, we can create three partitions, one for each consumer, the Kafka will spread the messages to each partition in a round robin fashion.

### Batching
Producers and consumers can push and pull one message at a time; however, we often push and pull messages in batches. For example, a consumer can grab a batch of ten messages at a time. This increases the latency of individual messages but typically improves network utilization, compression ratios, and throughput and is usually a worthwhile tradeoff. Kafka can even be configured to wait to accumulate messages, buffer up some work for until a set amount of time is reached, before processing.

### Push or Pull?
Messages are delivered in one of two ways: _push_-based or _pull_-based models.
- Push messages are sent via HTTP or gRPC. The message must be handled immediately and handle errors gracefully.
- Pull is effectively polling a message source over and over until a message is ready.

Producers usually push messages to a broker. Kafka does not allow pulling messages from producers into the broker. Consumers, on the other hand, can be either push or pull. Kafka prefers allowing consumers to pull messages when they are ready.

Kafka and other brokers push messages from producer to broker and pull messages from broker to consumer.

## References
- [Kafka Quickstart](https://kafka.apache.org/quickstart)
- [Kafka Design](https://kafka.apache.org/documentation/#majordesignelements)
- [Kafdrop](https://github.com/obsidiandynamics/kafdrop)