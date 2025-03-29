package repositories

import "Multidiciplinario/src/sensor_luz/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.LightSensorLDR) error
}
