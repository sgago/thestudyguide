# Apache Flink

Apache Flink is a streaming data processing pipeline. It is useful for scalable streaming ETL, analytics, and event-driven applications and services.

## Stream processing
Often, data come in as a stream: sensor voltages, stock prices, web events, and so on. These data can be organized into either a *bounded* or *unbounded* streams.

Say we have a numbers pouring in:
```
              |  Bounded  |   | Bounded |
past nums ... 1 2 3 4 5 6 7 8 9 10 ... future nums
              | <- Unbounded -> ...             |
                  | <- Another unbounded -> ... |
```

*Batch processing* is when we process a bounded stream, that is, when we have the entire dataset.
For example, computing a global statistics (max, min, average, etc.) of all numbers requires, well,
all the numbers to be present.

*Stream processing* is when we process numbers that may never end. For example, a sensor could detect
temperatures forever. Therefore, we need to process the numbers as they arrive.

In Flink, applications are composed of *streaming dataflows* that are transformed via *operators*.
These dataflows are directed graphs that begin from a data *source* and are output to some destination or *sink*.

```java
public static void main(String[] args) throws Exception {
    final StreamExecutionEnvironment env =
            StreamExecutionEnvironment.getExecutionEnvironment();

    // Source - where to get the data
    // Can be databases, other pipelines, flat files, etc.
    DataStream<Person> flintstones = env.fromElements(
            new Person("Fred", 35),
            new Person("Wilma", 35),
            new Person("Pebbles", 2));

    // Operator - does something to the data.
    // In the case it's a filter. We can map data, enrich, filter, combine,
    // group by, order it, etc.
    // Transformation - Does something to the data.
    // Sometimes, transformations consist of many operators.
    // Sometimes, transformations are one-to-one with operators but not always.
    DataStream<Person> adults = flintstones.filter(new FilterFunction<Person>() {
        @Override
        public boolean filter(Person person) throws Exception {
            return person.age >= 18;
        }
    });

    // Sink - a place to send the data
    // (more typically like operator.addSink(new MySink(...)))
    // Sinks can be databases, other streams, flat files, etc.
    adults.print();

    env.execute();
}
```

## Parallel dataflows
Flink programs are inherently parallel and distributed.
Streams are divided into *stream partitions*. Each stream partition
has one to many *operator subtasks*. Operator subtasks are independent
of one another and may operate on different treads or machines entirely.

The number of operator subtasks is given by the *parallelism* of that
operator. Operators may have different numbers of parallelism.
A sink task might only need one operator, but a calculation-heavy
operator could benefit from many.

*Streams* transport data between the operators that are on different
threads or machines. *One-to-one* is a forwarding pattern that preserves
partitioning and order of elements. Sources are typically
one-to-one. *Redistributing* streams, such as from the map, window, or apply operators change
the partitioning of streams. For example,
if we need to keyBy a session ID or similar, then the elements
will be redistributed as needed. This results in *non-determinism*;
there's no guarantee that elements will be in the same order
when we're sorting, grouping, and aggregating them.

## Timely stream processing
Typically, for many applications, we want deterministic, consistent results if we reprocess the data.
Therefore, we may want to pay attention
to when events occurred, when they are processed, or when they are completed. Elements have an event time timestamps,
processed timestamps, and ingestion timestamps, that we can use instead of wall-clock time of a particular machine.
There are some important differences on how these times are used, which will be covered later.

## Stateful stream processing
Operators in Flink can be stateful. An event can be handled differently  based on the accumulated effect of all events
before it. State can be used to count events per minute or something more complicated like fraud detection.

Since operators are on different threads and even different machines, state is typically in a sharded key-value store.
State is accessed locally via JVM heap or in other on-disk data structures.

## Fault tolerance via state snapshots
Flink can provide fault-tolerant, exactly-once semantics through state snapshots and stream replay.
Snapshots capture the entire state of the stream through the entire pipeline. If a failure occurs,
sources are rewound, state is restored, and processing is resumed. Snapshots are captured asynchronously to
avoid impeding element processing.

## Flink APIs
Now, there are different levels of abstraction for developing Flink applications. From the lowest level abstraction to
highest, they are:
1. Stateful stream processing - Exposed via the DataStream API's [Process function](https://nightlies.apache.org/flink/flink-docs-release-1.19/docs/dev/datastream/operators/process_function/). Here we can process events from streams freely and can register event and processing time callbacks.
2. DataStream / DataSet API - Many applications do not need the sophistication of the lowest level stateful stream processing. The DataStream API provides many common functions like transformations, joins, aggregations, windows, state, etc. Stateful stream processing can be integrated with the DataStream API.
3. Table API - Centered on tables centered on schemas and provides operators like projections, joins, group-bys, etc. It worries more about what logical operation should happen over specifying exactly how it should be done. This is, however, less expressive than the DataStream (core) APIs.
4. SQL - That's right. You can represent a Flink pipe via a SQL query expression. It uses the same Table API.

## What can be streamed?
- Primitives: String, Long, Integer, Boolean, Array
- Composites: Tuples, POJOs

Tuples from Tuple0 to Tuple25 can all be streamed:
```java
Tuple2<String, Integer> person = Tuple2.of("Fred", 35);

String name = person.f0;
Integer age = person.f1;
```

POJOs must be
- Public and standalone (no non-static inner class)
- Public empty constructor (no arguments)
- All non-static, non-transient fields are public, non-final. Alternatively, they have public getter- and setter-methods that follow the Java Beans naming conventions.

For example:
```java
public static class Person {
    public String name;
    public Integer age;
    public Person() {}

    public Person(String name, Integer age) {
        this.name = name;
        this.age = age;
    }

    public String toString() {
        return this.name + ": age " + this.age.toString();
    }
}
```

## Stream execution environment
Every Flink application needs an execution environment, typically called `env` or similar.
Streaming applications needs the `StreamExecutionEnvironment`. API calls to the `StreamExecutionEnvironment`
will build a job graph. `env.Execute()` sends the graph to the Flink *JobManager* which will parallelizes the job
and distributes slices of the job to *TaskManagers* for execution. Each parallel job slice will take up a *task slot*.

Note that the runtime requires the application to be serializable and all dependencies must be available to each
node on the cluster.

## Stream sources
Sources are inputs to the pipeline; sources can be in-memory collections, databases, flat-files, or even other pipelines.
Note: you may need to go grab appropriate dependencies.

```java
// In memory example
List<Person> people = new ArrayList<Person>();

people.add(new Person("Fred", 35));
people.add(new Person("Wilma", 35));
people.add(new Person("Pebbles", 2));

DataStream<Person> flintstones = env.fromCollection(people);

// From sockets
DataStream<String> lines = env.socketTextStream("localhost", 9999);
// Forward a port with netcat via `nc -lk 9999` and
// enter whatever text you want into the netcat terminal window

// From a file. Multiple file types are supported:
// CSV, NDJSON, Avro, Parquet, Orc, Debezium JSON, raw.
final FileSource<String> source = FileSource.forRecordStreamFormat(
    new TextLineInputFormat(),
    new Path("path/to/your/file.txt")
).build();

// From MongoDB, a more complex example
MongoSource<String> source = MongoSource.<String>builder()
    .setUri("mongodb://user:password@127.0.0.1:27017")
    .setDatabase("my_db")
    .setCollection("my_coll")
    .setProjectedFields("_id", "f0", "f1")
    .setFetchSize(2048)
    .setLimit(10000)
    .setNoCursorTimeout(true)
    .setPartitionStrategy(PartitionStrategy.SAMPLE)
    .setPartitionSize(MemorySize.ofMebiBytes(64))
    .setSamplesPerPartition(10)
    .setDeserializationSchema(new MongoDeserializationSchema<String>() {
        @Override
        public String deserialize(BsonDocument document) {
            return document.toJson();
        }

        @Override
        public TypeInformation<String> getProducedType() {
            return BasicTypeInfo.STRING_TYPE_INFO;
        }
    })
    .build();
```

Sources can be SQL DBs, NoSQL DBs, REST endpoints, or message brokers.
Some assembly might be required.

## Stream sinks
Pipeline sinks, a.k.a. outputs of the pipeline, can also be
files, message queues, databases, and other pipelines.
Note: you may need to go grab appropriate dependencies.

A file sink, for instance:
```java
// A "simple" file sink
final FileSink<String> sink = FileSink
    .forRowFormat(
        new Path(outputPath),
        new SimpleStringEncoder<String>("UTF-8"))
    .withRollingPolicy(
        DefaultRollingPolicy.builder()
            .withRolloverInterval(Duration.ofMinutes(15))
            .withInactivityInterval(Duration.ofMinutes(5))
            .withMaxPartSize(MemorySize.ofMebiBytes(1024))
            .build())
    .build();

someDataStream.sinkTo(sink);

// A MongoDB sink:
MongoSink<String> sink = MongoSink.<String>builder()
    .setUri("mongodb://user:password@127.0.0.1:27017")
    .setDatabase("my_db")
    .setCollection("my_coll")
    .setBatchSize(1000)
    .setBatchIntervalMs(1000)
    .setMaxRetries(3)
    .setDeliveryGuarantee(DeliveryGuarantee.AT_LEAST_ONCE)
    .setSerializationSchema(
        (input, context) -> new InsertOneModel<>(BsonDocument.parse(input)))
    .build();

someDataStream.sinkTo(sink);
```

## Operations


## References
1. https://nightlies.apache.org/flink/flink-docs-release-1.19/