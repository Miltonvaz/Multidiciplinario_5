package domain

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
)

type IBodyTemperature interface {
	Save(entities.BodyTemperature) (entities.BodyTemperature, error)
	GetMeasurementByID(id, userID int) (entities.BodyTemperature, error)
	GetLatestMeasurement(userID int) (entities.BodyTemperature, error)
	GetAllMeasurements(userID int) ([]entities.BodyTemperature, error)
	Delete(id, userID int) error
	GetAverageTemperature(userID int) (float64, error)
}
