package use_case

import "Multidiciplinario/src/sensor_luz/domain"

type GetMeasurementByID struct {
	db domain.ILightSensor
}

func NewGetMeasurementByID(db domain.ILightSensor) *GetMeasurementByID {
	return &GetMeasurementByID{db: db}
}

func (u *GetMeasurementByID) Execute(id int) (interface{}, error) {
	return u.db.GetMeasurementByID(id)
}
