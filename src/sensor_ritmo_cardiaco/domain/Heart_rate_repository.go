package domain

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
)

type IHeartRate interface {
	Save(entities.HeartRate) (entities.HeartRate, error)
	GetMeasurementByID(id int) (entities.HeartRate, error)
	GetLatestMeasurement() (entities.HeartRate, error)
	GetAllMeasurements() ([]entities.HeartRate, error)
	Delete(id int) error
	GetAverageHeartRate() (float64, error)
}
