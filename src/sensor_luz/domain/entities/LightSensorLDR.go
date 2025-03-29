package entities

type LightSensorLDR struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Lux    float64 `json:"luz"`
}
