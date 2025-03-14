package use_case

import "Multidiciplinario/src/sensor_temperatura_humedad/domain"

type GetAverageTemperature struct {
	db domain.ITemperatureAndHumidity
}

func NewGetAverageTemperature(db domain.ITemperatureAndHumidity) *GetAverageTemperature {
	return &GetAverageTemperature{db: db}
}

func (gt *GetAverageTemperature) Execute() (float64, error) {
	return gt.db.GetAverageTemperature()
}
