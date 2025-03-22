package repositories

import "Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.HeartRate) error
}
