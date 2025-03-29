package use_case

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/domain"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
)

type Get_Latest struct {
	db domain.IBodyTemperature
}

func NewGet_Latest(db domain.IBodyTemperature) *Get_Latest {
	return &Get_Latest{db: db}
}

func (gt *Get_Latest) Execute(userID int) (entities.BodyTemperature, error) {
	return gt.db.GetLatestMeasurement(userID)
}
