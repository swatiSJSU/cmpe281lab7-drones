package service

import dronescommon "github.com/cloudnativego/drones-common"

type telemetryEvent struct {
	DroneID          string `json:"drone_id"`
	RemainingBattery int    `json:"battery"`
	Uptime           int    `json:"uptime"`
	CoreTemp         int    `json:"core_temp"`
	ReceivedOn       string `json:"received_on"`
}

type alertEvent struct {
	DroneID     string `json:"drone_id"`
	FaultCode   int    `json:"fault_code"`
	Description string `json:"description"`
	ReceivedOn  string `json:"received_on"`
}

type positionEvent struct {
	DroneID         string  `json:"drone_id"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Altitude        float32 `json:"altitude"`
	CurrentSpeed    float32 `json:"current_speed"`
	HeadingCardinal int     `json:"heading_cardinal"`
	ReceivedOn      string  `json:"received_on"`
}

type eventRepository interface {
	UpdateLastTelemetryEvent(telemetryEvent dronescommon.TelemetryUpdatedEvent) (err error)
	UpdateLastAlertEvent(alertEvent dronescommon.AlertSignalledEvent) (err error)
	UpdateLastPositionEvent(positionEvent dronescommon.PositionChangedEvent) (err error)
	GetTelemetryEvent(droneID string) (event dronescommon.TelemetryUpdatedEvent, err error)
	GetPositionEvent(droneID string) (event dronescommon.PositionChangedEvent, err error)
	GetAlertEvent(droneID string) (event dronescommon.AlertSignalledEvent, err error)
}
