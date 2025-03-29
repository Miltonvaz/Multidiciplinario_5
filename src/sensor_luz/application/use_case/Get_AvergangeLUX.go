package use_case

import "Multidiciplinario/src/sensor_luz/domain"

type GetAverageLightLDR struct {
	db domain.ILightSensor
}

func NewGetAverageLightLDR(db domain.ILightSensor) *GetAverageLightLDR {
	return &GetAverageLightLDR{db: db}
}

func (gt *GetAverageLightLDR) Execute(userID int) (float64, error) {
	return gt.db.GetAverageLux(userID)
}
