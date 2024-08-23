```zsh
curl -X POST http://localhost:8081/add_node \
     -H "Content-Type: application/json" \
     -d '{"node_id": "goginair-3", "address": "goginair-3:3001"}'

curl -X POST http://localhost:8082/add_node \
     -H "Content-Type: application/json" \
     -d '{"node_id": "goginair-1", "address": "goginair-1:3001"}'

curl -X POST http://localhost:8082/add_node \
     -H "Content-Type: application/json" \
     -d '{"node_id": "goginair-2", "address": "goginair-2:3001"}'


curl -X POST -i -H 'Content-Type: application/json' -H 'User-Agent: go-resty/2.14.0 (https://github.com/go-resty/resty)' -d '{"value":"your_value_here2"}' http://192.168.0.4:3001/set_value

curl -X POST -i -H 'Content-Type: application/json' -d '{"value":"your_value_here2"}' http://goginair-2:3001/set_value
```