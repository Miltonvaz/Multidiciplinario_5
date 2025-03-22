package entities

type GyroscopeSensor struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	XAxis  float64 `json:"giroX"`
	YAxis  float64 `json:"giroY"`
	ZAxis  float64 `json:"giroZ"`
}
