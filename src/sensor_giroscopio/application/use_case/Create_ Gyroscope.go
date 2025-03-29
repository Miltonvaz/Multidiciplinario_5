package use_case

import (
	"Multidiciplinario/src/sensor_giroscopio/application/repositories"
	"Multidiciplinario/src/sensor_giroscopio/domain"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"
	"log"
)

type Create_GyroscopeSensor struct {
	gyroscopeRepo       domain.IGyroscopeSensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_GyroscopeSensor(gyroscopeRepo domain.IGyroscopeSensor, serviceNotification *repositories.ServiceNotification) *Create_GyroscopeSensor {
	return &Create_GyroscopeSensor{
		gyroscopeRepo:       gyroscopeRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_GyroscopeSensor) Execute(gyroscope entities.GyroscopeSensor) (entities.GyroscopeSensor, error) {
	created, err := c.gyroscopeRepo.Save(gyroscope)

	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notificando medición creada: %v", err)
		return entities.GyroscopeSensor{}, err
	}

	return created, nil
}
