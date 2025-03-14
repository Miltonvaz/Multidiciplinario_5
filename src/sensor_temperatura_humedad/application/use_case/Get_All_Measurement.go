package use_case

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/domain"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"
)

type Get_All_Measurement struct {
	db domain.ITemperatureAndHumidity
}

func NewGet_All(db domain.ITemperatureAndHumidity) *Get_All_Measurement {
	return &Get_All_Measurement{db: db}
}

func (gt *Get_All_Measurement) Execute() ([]entities.TemperatureAndHumidity, error) {
	return gt.db.GetAllMeasurements()
}
