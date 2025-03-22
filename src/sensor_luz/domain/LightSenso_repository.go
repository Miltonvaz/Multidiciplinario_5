package domain

import "Multidiciplinario/src/sensor_luz/domain/entities"

type ILightSensor interface {
	Save(entities.LightSensorLDR) (entities.LightSensorLDR, error)
	GetMeasurementByID(id int) (entities.LightSensorLDR, error)
	GetLatestMeasurement() (entities.LightSensorLDR, error)
	GetAllMeasurements() ([]entities.LightSensorLDR, error)
	Delete(id int) error
	GetAverageLux() (float64, error)
}
