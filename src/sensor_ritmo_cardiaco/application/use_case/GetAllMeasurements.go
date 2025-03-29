package use_case

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
)

type Get_All_Measurement struct {
	db domain.IHeartRate
}

func NewGet_All(db domain.IHeartRate) *Get_All_Measurement {
	return &Get_All_Measurement{db: db}
}

func (gt *Get_All_Measurement) Execute(userID int) ([]entities.HeartRate, error) {
	return gt.db.GetAllMeasurements(userID)
}
