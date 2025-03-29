package use_case

import "Multidiciplinario/src/sensor_giroscopio/domain"

type GetGyroscopeSensorByID struct {
	db domain.IGyroscopeSensor
}

func NewGetGyroscopeSensorByID(db domain.IGyroscopeSensor) *GetGyroscopeSensorByID {
	return &GetGyroscopeSensorByID{db: db}
}

func (u *GetGyroscopeSensorByID) Execute(id, userID int) (interface{}, error) {
	return u.db.GetMeasurementByID(id, userID)
}
