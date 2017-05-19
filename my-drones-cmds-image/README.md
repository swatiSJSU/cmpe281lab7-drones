[![wercker status](https://app.wercker.com/status/0e2174b4304cf85753b912cc4ca0aafe/m/master "wercker status")](https://app.wercker.com/project/bykey/0e2174b4304cf85753b912cc4ca0aafe)

# Drone Army Sample - Command Processing Service
This is part of the Drone Army sample suite, this service is responsible for processing incoming commands
and converting them into events for dispatch into queues.

## RESTful Endpoints
The following is a list of the REST endpoints exposed by this service.

| Resource | Method | Description |
|---|---|---|
| /api/cmds/telemetry | POST | Submits a new telemetry update command, dispatches corresponding event to queue |
| /api/cmds/alerts | POST | Submits a new alert command |
| /api/cmds/position | POST | Submits a new position update command |
