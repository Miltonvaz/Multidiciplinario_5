package entities

type BodyTemperature struct {
	ID          int     `json:"id"`
	Temperature float64 `json:"temperatura"`
	UserID      int     `json:"user_id"`
}
