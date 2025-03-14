package use_case

import "Multidiciplinario/src/sensor_temperatura_humedad/domain"

type GetMeasurementByID struct {
	db domain.ITemperatureAndHumidity
}

func NewGetMeasurementByID(db domain.ITemperatureAndHumidity) *GetMeasurementByID {
	return &GetMeasurementByID{db: db}
}

func (u *GetMeasurementByID) Execute(id int) (interface{}, error) {
	return u.db.GetMeasurementByID(id)
}
