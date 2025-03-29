package use_case

import "Multidiciplinario/src/sensor_ritmo_cardiaco/domain"

type Delete_HeartRate struct {
	db domain.IHeartRate
}

func NewDelete_HeartRate(db domain.IHeartRate) *Delete_HeartRate {
	return &Delete_HeartRate{db: db}
}

func (dt *Delete_HeartRate) Execute(id int, userID int) error {
	return dt.db.Delete(id, userID)
}
