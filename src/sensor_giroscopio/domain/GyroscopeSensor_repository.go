package domain

import entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"

type IGyroscopeSensor interface {
	Save(entities.GyroscopeSensor) (entities.GyroscopeSensor, error)
	GetMeasurementByID(id int) (entities.GyroscopeSensor, error)
	GetLatestMeasurement() (entities.GyroscopeSensor, error)
	GetAllMeasurements() ([]entities.GyroscopeSensor, error)
	Delete(id int) error
	GetAverageGyroscopeData() (float64, error)
}
