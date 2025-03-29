package domain

import "Multidiciplinario/src/sensor_luz/domain/entities"

type ILightSensor interface {
	Save(entities.LightSensorLDR) (entities.LightSensorLDR, error)
	GetMeasurementByID(id int, userID int) (entities.LightSensorLDR, error)
	GetLatestMeasurement(userID int) (entities.LightSensorLDR, error)
	GetAllMeasurements(userID int) ([]entities.LightSensorLDR, error)
	Delete(id int, userID int) error
	GetAverageLux(userID int) (float64, error)
}
