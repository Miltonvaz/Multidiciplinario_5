package domain

import "Multidiciplinario/src/sensor_luz/domain/entities"

type ILightSensor interface {
	Save(measurement entities.LightSensorLDR) error
	GetMeasurementByID(id int) (entities.LightSensorLDR, error)
	GetLatestMeasurement(deviceID string) (entities.LightSensorLDR, error)
	GetAllMeasurements() ([]entities.LightSensorLDR, error)
	Delete(id int) error
	GetAverageLux(deviceID string) (float64, error)
}
