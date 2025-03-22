package entities

type GyroscopeSensor struct {
	ID    int     `json:"id"`
	XAxis float64 `json:"x_axis"`
	YAxis float64 `json:"y_axis"`
	ZAxis float64 `json:"z_axis"`
}
