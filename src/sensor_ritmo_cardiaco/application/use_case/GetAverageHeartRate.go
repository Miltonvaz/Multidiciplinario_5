package use_case

import "Multidiciplinario/src/sensor_ritmo_cardiaco/domain"

type GetAverageHeartRate struct {
	db domain.IHeartRate
}

func NewGetAverageHeartRate(db domain.IHeartRate) *GetAverageHeartRate {
	return &GetAverageHeartRate{db: db}
}

func (gt *GetAverageHeartRate) Execute() (float64, error) {
	return gt.db.GetAverageHeartRate()
}
