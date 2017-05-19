# pull mongo
```
docker pull mongo:latest
```

# run mongo
```
docker run --name my-mongo -d -p 27017:27017 mongo
```

# Build modified drone-cmds image
```
docker build -t my-drone-events-image .
```

# Run drone-cmds and link to rabbitmq and mongo
```
# Assuming that rabbitmq is already running from the drone-cmds service

docker run -d --name my-drone-events-app \
              --link my-rabbit:rabbit \
              --link my-mongo:mongo \
              -p 3001:3000 \
              my-drone-events-image
```

- Any command inserted into the drone-cmds should be rolled up into the mongodb
```
curl -i -X POST http://localhost:3000/api/cmds/telemetry \
     -H "Content-Type: application/json" \
     -d "{\"drone_id\":\"drone999\",\"battery\":72,\"uptime\":6941,\"core_temp\":21}"
```
- Check the mongodb server to verify if the service is working fine






