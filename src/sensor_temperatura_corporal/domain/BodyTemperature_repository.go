package domain

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
)

type IBodyTemperature interface {
	Save(entities.BodyTemperature) (entities.BodyTemperature, error)
	GetMeasurementByID(id int) (entities.BodyTemperature, error)
	GetLatestMeasurement() (entities.BodyTemperature, error)
	GetAllMeasurements() ([]entities.BodyTemperature, error)
	Delete(id int) error
	GetAverageTemperature() (float64, error)
}
