# Outbox

Imagine you need to store products into a database and a search engine. A basic implementation sends product adds, updates, and deletes to both places. But now, imagine, there's a network issue and database doesn't get updated. Now our database and search engine diverge; customers can search for a product but not get details or buy it!

An outbox pattern is used to durably replicate data to different data stores. For example, say that we want a SQL database to store product data in, but now we also want a search engine for quick look ups on our website. Now, assuming that we're ok with sequential consistency between the two, we can write database changes to a dedicated outbox table. This outbox table will list all the products that are created, updated, and deleted and a separate relay process will write these changes into our search engine.

## Debezium
Fortunately, we don't even need to write most of this oursevles. [Debezium](https://debezium.io/documentation/reference/stable/tutorial.html#introduction-debezium) can capture table changes and write them into a Kafka stream for us to update the search engine. Pretty sweet, right?

```zsh
psql -h localhost -p 5432 -U postgres -d postgres
```

```
\c mycompany
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