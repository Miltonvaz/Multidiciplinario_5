package domain

import entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"

type IGyroscopeSensor interface {
	Save(entities.GyroscopeSensor) (entities.GyroscopeSensor, error)
	GetMeasurementByID(id int, userID int) (entities.GyroscopeSensor, error)
	GetLatestMeasurement(userID int) (entities.GyroscopeSensor, error)
	GetAllMeasurements(userID int) ([]entities.GyroscopeSensor, error)
	Delete(id int, userID int) error
	GetAverageGyroscopeData(userID int) (float64, error)
}
