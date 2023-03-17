# Request counter

This is a simple exercise API which counts requests per server instance and outputs this count as well as the total across the cluster.

## Running

Build the image:

```
docker build -t counter .
```

Then init the swarm:

```
docker swarm init
```

Finally, deploy the swarm:

```
docker stack deploy --compose-file docker-compose.yml counter
```

This will start a redis instance on port `6379` and start the API on port `8083`.

A `GET` request to `http://localhost:8083/` will output the counter stats.