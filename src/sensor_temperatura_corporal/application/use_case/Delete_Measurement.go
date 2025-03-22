package use_case

import "Multidiciplinario/src/sensor_temperatura_corporal/domain"

type Delete_BodyTemperature struct {
	db domain.IBodyTemperature
}

func NewDelete_TemperatureAndHumidity(db domain.IBodyTemperature) *Delete_BodyTemperature {
	return &Delete_BodyTemperature{db: db}
}

func (dt *Delete_BodyTemperature) Execute(id int) error {
	return dt.db.Delete(id)
}
