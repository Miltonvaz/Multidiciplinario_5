package entities

import "time"

type LightSensorLDR struct {
	ID        int       `json:"id"`
	Lux       float64   `json:"lux"`
	Timestamp time.Time `json:"timestamp"`
}
