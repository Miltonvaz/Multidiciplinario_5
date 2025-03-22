package use_case

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/domain"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
)

type Get_All_Measurement struct {
	db domain.IBodyTemperature
}

func NewGet_All(db domain.IBodyTemperature) *Get_All_Measurement {
	return &Get_All_Measurement{db: db}
}

func (gt *Get_All_Measurement) Execute() ([]entities.BodyTemperature, error) {
	return gt.db.GetAllMeasurements()
}
