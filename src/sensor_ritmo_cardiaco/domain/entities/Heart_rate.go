package entities

type HeartRate struct {
	ID        int     `json:"id"`
	HeartRate float64 `json:"bpm"`
	UserID    int     `json:"user_id"`
}
