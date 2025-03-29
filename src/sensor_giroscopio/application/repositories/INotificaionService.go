package repositories

import entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.GyroscopeSensor) error
}
