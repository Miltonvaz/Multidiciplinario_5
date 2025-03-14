package domain

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"
)

type ITemperatureAndHumidity interface {
	Save(entities.TemperatureAndHumidity) (entities.TemperatureAndHumidity, error)
	GetMeasurementByID(id int) (entities.TemperatureAndHumidity, error)
	GetLatestMeasurement() (entities.TemperatureAndHumidity, error)
	GetAllMeasurements() ([]entities.TemperatureAndHumidity, error)
	Delete(id int) error
	GetAverageTemperature() (float64, error)
	GetAverageHumidity() (float64, error)
}
