package use_case

import "Multidiciplinario/src/sensor_temperatura_humedad/domain"

type GetAverageHumidity struct {
	db domain.ITemperatureAndHumidity
}

func NewGetAverageHumidity(db domain.ITemperatureAndHumidity) *GetAverageHumidity {
	return &GetAverageHumidity{db: db}
}

func (gt *GetAverageHumidity) Execute() (float64, error) {
	return gt.db.GetAverageHumidity()
}
