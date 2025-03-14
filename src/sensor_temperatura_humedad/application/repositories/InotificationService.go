package repositories

import "Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.TemperatureAndHumidity) error
}
