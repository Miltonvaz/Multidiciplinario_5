package use_case

import "Multidiciplinario/src/sensor_temperatura_corporal/domain"

type GetMeasurementByID struct {
	db domain.IBodyTemperature
}

func NewGetMeasurementByID(db domain.IBodyTemperature) *GetMeasurementByID {
	return &GetMeasurementByID{db: db}
}

func (u *GetMeasurementByID) Execute(id, userID int) (interface{}, error) {
	return u.db.GetMeasurementByID(id, userID)
}
