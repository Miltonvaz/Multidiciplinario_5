package use_case

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/repositories"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
	"log"
)

type Create_BodyTemperature struct {
	appointmentRepo     domain.IBodyTemperature
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_BodyTemperature(appointmentRepo domain.IBodyTemperature, serviceNotification *repositories.ServiceNotification) *Create_BodyTemperature {
	return &Create_BodyTemperature{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_BodyTemperature) Execute(appointment entities.BodyTemperature) (entities.BodyTemperature, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return entities.BodyTemperature{}, err
	}
	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notifying about created appointment: %v", err)
		return entities.BodyTemperature{}, err
	}

	return created, nil
}
