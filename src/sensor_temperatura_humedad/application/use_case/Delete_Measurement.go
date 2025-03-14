package use_case

import "Multidiciplinario/src/sensor_temperatura_humedad/domain"

type Delete_TemperatureAndHumidity struct {
	db domain.ITemperatureAndHumidity
}

func NewDelete_TemperatureAndHumidity(db domain.ITemperatureAndHumidity) *Delete_TemperatureAndHumidity {
	return &Delete_TemperatureAndHumidity{db: db}
}

func (dt *Delete_TemperatureAndHumidity) Execute(id int) error {
	return dt.db.Delete(id)
}
