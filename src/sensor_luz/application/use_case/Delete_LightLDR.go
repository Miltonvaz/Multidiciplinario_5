package use_case

import "Multidiciplinario/src/sensor_luz/domain"

type Delete_LightLDR struct {
	db domain.ILightSensor
}

func NewDelete_LightLDR(db domain.ILightSensor) *Delete_LightLDR {
	return &Delete_LightLDR{db: db}
}

func (dt *Delete_LightLDR) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
