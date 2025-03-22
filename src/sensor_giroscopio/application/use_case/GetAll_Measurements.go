package use_case

import (
	"Multidiciplinario/src/sensor_giroscopio/domain"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"
)

type GetAllGyroscopeSensors struct {
	db domain.IGyroscopeSensor
}

func NewGetAllGyroscopeSensors(db domain.IGyroscopeSensor) *GetAllGyroscopeSensors {
	return &GetAllGyroscopeSensors{db: db}
}

func (gt *GetAllGyroscopeSensors) Execute() ([]entities.GyroscopeSensor, error) {
	return gt.db.GetAllMeasurements()
}
