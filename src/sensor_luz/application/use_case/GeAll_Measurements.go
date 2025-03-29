package use_case

import (
	"Multidiciplinario/src/sensor_luz/domain"
	"Multidiciplinario/src/sensor_luz/domain/entities"
)

type Get_All_Measurement struct {
	db domain.ILightSensor
}

func NewGet_All(db domain.ILightSensor) *Get_All_Measurement {
	return &Get_All_Measurement{db: db}
}

func (gt *Get_All_Measurement) Execute(userID int) ([]entities.LightSensorLDR, error) {
	return gt.db.GetAllMeasurements(userID)
}
