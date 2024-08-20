# Outbox
Imagine you need to store products into a database and a search engine. A basic implementation sends product adds, updates, and deletes to both places. But now, imagine, there's a network issue and database doesn't get updated. Now our database and search engine diverge; customers can search for a product but they can't buy it!

An outbox pattern is used to replicate data to different data stores. Assuming that we're ok with sequential consistency between the two, we can write DB table changes to a dedicated outbox table. Put another way, any time our products are added, updated, or deleted, the changes will be reflected into an outbox table. Then, a separate relay process will write these product changes into our search engine.

## Debezium
Fortunately, we don't even need to write the relay process ourselves. [Debezium](https://debezium.io/documentation/reference/stable/tutorial.html#introduction-debezium) can capture table changes and write them into a Kafka stream. In turn, Kafka can load the product updates into a search engine. Pretty sweet, right?

This project has several images that depend on one another. Debezium needs Kafka which needs Zookeeper. To show the outbox pattern in all its glory, it's nice to have a producer and consumer or PostgresSQL and ElasticSearch, respectively. In turn, a UI is helpful so we install Kafdrop and Kibana for Kafka and ElasticSearch, respectively. Lots of images to show this one off!

```zsh
psql -h localhost -p 5432 -U postgres -d postgres
```

```
\c outbox
```

```sql
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```sql
CREATE PUBLICATION product_publication
    FOR TABLE products
    WITH (publish = 'insert, update, delete');
```

Register the Postgres connector.
```zsh
curl -X POST -H "Content-Type: application/json" \
     -d @register-postgres.json \
     http://localhost:8083/connectors
```

And then register the ElasticSearch connector.
```zsh
curl -X POST -H "Content-Type: application/json" \
     -d @register-elasticsearch.json \
     http://localhost:8083/connectors
```

```sql
INSERT INTO products (name, description, price)
VALUES ('Sample Product', 'This is a sample product description.', 19.99);
```

## Troubleshooting

### Debezium
curl -X GET http://localhost:8083/connectors/
curl -X GET http://localhost:8083/connectors/product-connector/status
curl -X GET http://localhost:8083/connectors/elasticsearch-sink-connector/status

### Kafka
docker exec -it <KAFKA-CONTAINER-ID> kafka-topics --bootstrap-server localhost:9092 --list

### ElasticSearch
curl -X GET "localhost:9200/_cat/indices?v"

## References
- [Elasticsearch Sink Connector](https://d2p6pa21dvn84.cloudfront.net/api/plugins/confluentinc/kafka-connect-elasticsearch/versions/14.1.1/confluentinc-kafka-connect-elasticsearch-14.1.1.zip)