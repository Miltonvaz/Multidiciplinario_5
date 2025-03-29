package use_case

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/repositories"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"

	"log"
)

type Create_HeartRate struct {
	appointmentRepo     domain.IHeartRate
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_HeartRate(appointmentRepo domain.IHeartRate, serviceNotification *repositories.ServiceNotification) *Create_HeartRate {
	return &Create_HeartRate{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_HeartRate) Execute(appointment entities.HeartRate) (entities.HeartRate, error) {

	created, err := c.appointmentRepo.Save(appointment)

	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notificando cita creada: %v", err)
		return entities.HeartRate{}, err
	}

	return created, nil
}
