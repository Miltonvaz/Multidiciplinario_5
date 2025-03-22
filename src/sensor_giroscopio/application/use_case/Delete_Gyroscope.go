package use_case

import "Multidiciplinario/src/sensor_giroscopio/domain"

type Delete_GyroscopeSensor struct {
	db domain.IGyroscopeSensor
}

func NewDelete_GyroscopeSensor(db domain.IGyroscopeSensor) *Delete_GyroscopeSensor {
	return &Delete_GyroscopeSensor{db: db}
}

func (dt *Delete_GyroscopeSensor) Execute(id int) error {
	return dt.db.Delete(id)
}
