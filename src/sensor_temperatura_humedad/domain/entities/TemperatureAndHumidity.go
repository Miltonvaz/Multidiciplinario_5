package entities

type TemperatureAndHumidity struct {
	ID          int     `json:"id"`
	Temperature float64 `json:"temperatura"`
	Humidity    float64 `json:"humedad"`
}
