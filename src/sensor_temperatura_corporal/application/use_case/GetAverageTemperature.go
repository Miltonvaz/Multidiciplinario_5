package use_case

import "Multidiciplinario/src/sensor_temperatura_corporal/domain"

type GetAverageTemperature struct {
	db domain.IBodyTemperature
}

func NewGetAverageTemperature(db domain.IBodyTemperature) *GetAverageTemperature {
	return &GetAverageTemperature{db: db}
}

func (gt *GetAverageTemperature) Execute(userID int) (float64, error) {
	return gt.db.GetAverageTemperature(userID)
}
