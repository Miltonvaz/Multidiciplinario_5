package use_case

import (
	"Multidiciplinario/src/sensor_giroscopio/domain"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"
)

type Get_Latest_Gyroscope struct {
	db domain.IGyroscopeSensor
}

func NewGet_Latest_Gyroscope(db domain.IGyroscopeSensor) *Get_Latest_Gyroscope {
	return &Get_Latest_Gyroscope{db: db}
}

func (gt *Get_Latest_Gyroscope) Execute(userID int) (entities.GyroscopeSensor, error) {
	return gt.db.GetLatestMeasurement(userID)
}
