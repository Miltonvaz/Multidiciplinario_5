package use_case

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
)

type Get_Latest struct {
	db domain.IHeartRate
}

func NewGet_Latest(db domain.IHeartRate) *Get_Latest {
	return &Get_Latest{db: db}
}

func (gt *Get_Latest) Execute() (entities.HeartRate, error) {
	return gt.db.GetLatestMeasurement()
}
