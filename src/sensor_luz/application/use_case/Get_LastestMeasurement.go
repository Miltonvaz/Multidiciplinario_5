package use_case

import (
	"Multidiciplinario/src/sensor_luz/domain"
	"Multidiciplinario/src/sensor_luz/domain/entities"
)

type Get_Latest struct {
	db domain.ILightSensor
}

func NewGet_Latest(db domain.ILightSensor) *Get_Latest {
	return &Get_Latest{db: db}
}

func (gt *Get_Latest) Execute() (entities.LightSensorLDR, error) {
	return gt.db.GetLatestMeasurement()
}
