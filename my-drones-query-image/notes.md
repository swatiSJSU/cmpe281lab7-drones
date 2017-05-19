# Drones-query
### Build modified docker image
```
docker build -t my-drone-query-image .
```
# Run drones-query
```
docker run -d --name my-drone-query-app \
              --link my-rabbit:rabbit \
              --link my-mongo:mongo \
              -p 3002:3000 \
              my-drone-query-image

curl -i -X GET http://localhost:3002/drones/drone999/lastTelemetry
```

# Kong setup
drone-cmds api is listening  on localhost:3000  
drone-events service would be listening on localhost:3001  
and drone-query api would be listenting on localhost:3002  

We could place kong api gateway in front of these services to route to the
appropriate service

```
  # Run cassandra
	docker run -d --name kong-database -p 9042:9042 cassandra:2.2

  # Run kong. Link to cassandra and drone services
	docker run -d --name kong \
              --link my-drone-cmds-app:drone-cmds-app \
              --link my-drone-query-app:drone-query-app \
              --link kong-database:kong-database \
              -e "KONG_DATABASE=cassandra" \
              -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
              -e "KONG_PG_HOST=kong-database" \
              -p 8000:8000 \
              -p 8443:8443 \
              -p 8001:8001 \
              -p 7946:7946 \
              -p 7946:7946/udp \
              kong
```

Kong gateway is listening on port 8000  
Kong adming api is listenting on port 8001  

# Adding apis to kong by sending requests to port 8001
```
  # Add route for the command api
  curl -i -X POST \
  --url http://localhost:8001/apis/ \
  --data 'name=command' \
  --data 'upstream_url=http://drone-cmds-app:3000' \
  --data 'uris=/command'

  # Add route for the query api
  curl -i -X POST \
  --url http://localhost:8001/apis/ \
  --data 'name=query' \
  --data 'upstream_url=http://drone-query-app:3000' \
  --data 'uris=/query'

  # Test if the routes got added successfully
  curl -i -X GET http://localhost:8001/apis
```

# Testing the whole setup
```
curl -i -X POST http://localhost:8000/command/api/cmds/telemetry \
     -H "Content-Type: application/json" \
     -d "{\"drone_id\":\"drone100\",\"battery\":22,\"uptime\":6900,\"core_temp\":21}"

curl -i -X GET http://localhost:8000/query/drones/drone100/lastTelemetry
```

