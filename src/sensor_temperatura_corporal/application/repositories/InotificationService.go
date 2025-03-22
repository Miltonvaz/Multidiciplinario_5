package repositories

import "Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.BodyTemperature) error
}
