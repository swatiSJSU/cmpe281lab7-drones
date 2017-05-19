[![wercker status](https://app.wercker.com/status/51c65e97fa901561dc1720e1ed598106/m "wercker status")](https://app.wercker.com/project/bykey/51c65e97fa901561dc1720e1ed598106)

# Drone Army Query Service
Service for the Event Sourcing/CQRS sample that fulfills the role of query service. Exposes a simple API for
querying the last event stored for any given drone ID.

| Resource | Method | Description |
|---|---|---|
| /drones/{droneId}/lastTelemetry | GET | Retrieves the last telemetry event for a drone |
| /drones/{droneId}/lastAlert | GET | Retrieves the last alert event for a drone |
| /drones/{droneId}/lastPosition | GET | Retrieves the last position event for a drone |
