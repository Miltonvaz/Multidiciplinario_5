package use_case

import "Multidiciplinario/src/sensor_giroscopio/domain"

type GetAverageGyroscopeData struct {
	db domain.IGyroscopeSensor
}

func NewGetAverageGyroscopeData(db domain.IGyroscopeSensor) *GetAverageGyroscopeData {
	return &GetAverageGyroscopeData{db: db}
}

func (gt *GetAverageGyroscopeData) Execute(userID int) (float64, error) {
	return gt.db.GetAverageGyroscopeData(userID)
}
