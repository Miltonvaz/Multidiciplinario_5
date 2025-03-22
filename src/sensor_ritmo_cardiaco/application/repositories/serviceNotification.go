package repositories

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
	"log"
)

type ServiceNotification struct {
	notificationPort NotificationPort
}

func NewServiceNotification(notificationPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{notificationPort: notificationPort}
}

func (sn *ServiceNotification) NotifyAppointmentCreated(appointment entities.HeartRate) error {
	log.Println("Notificando la creación de la cita...")

	err := sn.notificationPort.PublishEvent("cita_creada", appointment)
	if err != nil {
		log.Printf("Error al publicar el evento: %v", err)
		return err
	}
	return nil
}
