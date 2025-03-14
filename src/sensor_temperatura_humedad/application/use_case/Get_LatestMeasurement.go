package use_case

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/domain"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"
)

type Get_Latest struct {
	db domain.ITemperatureAndHumidity
}

func NewGet_Latest(db domain.ITemperatureAndHumidity) *Get_Latest {
	return &Get_Latest{db: db}
}

func (gt *Get_Latest) Execute() (entities.TemperatureAndHumidity, error) {
	return gt.db.GetLatestMeasurement()
}
