package use_case

import "Multidiciplinario/src/sensor_ritmo_cardiaco/domain"

type GetMeasurementByID struct {
	db domain.IHeartRate
}

func NewGetMeasurementByID(db domain.IHeartRate) *GetMeasurementByID {
	return &GetMeasurementByID{db: db}
}

func (u *GetMeasurementByID) Execute(id int, userID int) (interface{}, error) {
	return u.db.GetMeasurementByID(id, userID)
}
