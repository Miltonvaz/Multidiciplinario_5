package domain

import "Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"

type IHeartRate interface {
	Save(entities.HeartRate) (entities.HeartRate, error)
	GetMeasurementByID(id int, userID int) (entities.HeartRate, error)
	GetLatestMeasurement(userID int) (entities.HeartRate, error)
	GetAllMeasurements(userID int) ([]entities.HeartRate, error)
	Delete(id int, userID int) error
	GetAverageHeartRate(userID int) (float64, error)
}
