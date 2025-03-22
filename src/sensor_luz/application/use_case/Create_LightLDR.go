package use_case

import (
	"Multidiciplinario/src/sensor_luz/application/repositories"
	"Multidiciplinario/src/sensor_luz/domain"
	"Multidiciplinario/src/sensor_luz/domain/entities"

	"log"
)

type Create_LightLDR struct {
	appointmentRepo     domain.ILightSensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_LightLDR(appointmentRepo domain.ILightSensor, serviceNotification *repositories.ServiceNotification) *Create_LightLDR {
	return &Create_LightLDR{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_LightLDR) Execute(appointment entities.LightSensorLDR) (entities.LightSensorLDR, error) {

	created, err := c.appointmentRepo.Save(appointment)

	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notificando cita creada: %v", err)
		return entities.LightSensorLDR{}, err
	}

	return created, nil
}
