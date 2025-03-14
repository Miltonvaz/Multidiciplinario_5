package use_case

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/application/repositories"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"
	"log"
)

type Create_TemperatureAndHumidity struct {
	appointmentRepo     domain.ITemperatureAndHumidity
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_TemperatureAndHumidity(appointmentRepo domain.ITemperatureAndHumidity, serviceNotification *repositories.ServiceNotification) *Create_TemperatureAndHumidity {
	return &Create_TemperatureAndHumidity{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_TemperatureAndHumidity) Execute(appointment entities.TemperatureAndHumidity) (entities.TemperatureAndHumidity, error) {

	created, err := c.appointmentRepo.Save(appointment)

	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notificando cita creada: %v", err)
		return entities.TemperatureAndHumidity{}, err
	}

	return created, nil
}
