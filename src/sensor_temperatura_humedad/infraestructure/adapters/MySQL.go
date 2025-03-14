package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"Multidiciplinario/src/sensor_temperatura_humedad/domain"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

// NewMySQL retorna una nueva instancia del adaptador MySQL que implementa la interfaz ITemperatureAndHumidity
func NewMySQL(conn *sql.DB) domain.ITemperatureAndHumidity {
	return &MySQL{conn: conn}
}

// Save guarda una nueva medición de temperatura y humedad
func (m *MySQL) Save(sensor entities.TemperatureAndHumidity) (entities.TemperatureAndHumidity, error) {
	query := `INSERT INTO temperature_and_humidity (temperature, humidity) VALUES (?, ?)`
	_, err := m.conn.Exec(query, sensor.Temperature, sensor.Humidity)
	if err != nil {
		return entities.TemperatureAndHumidity{}, err
	}
	return sensor, nil
}

// GetMeasurementByID obtiene una medición por su ID
func (m *MySQL) GetMeasurementByID(id int) (entities.TemperatureAndHumidity, error) {
	var sensor entities.TemperatureAndHumidity
	query := `SELECT id, temperature, humidity FROM temperature_and_humidity WHERE id = ?`
	err := m.conn.QueryRow(query, id).Scan(&sensor.ID, &sensor.Temperature, &sensor.Humidity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

// GetLatestMeasurement obtiene la última medición registrada
func (m *MySQL) GetLatestMeasurement() (entities.TemperatureAndHumidity, error) {
	var sensor entities.TemperatureAndHumidity
	query := `SELECT id, temperature, humidity FROM temperature_and_humidity ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query).Scan(&sensor.ID, &sensor.Temperature, &sensor.Humidity)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

// GetAllMeasurements obtiene todas las mediciones
func (m *MySQL) GetAllMeasurements() ([]entities.TemperatureAndHumidity, error) {
	query := "SELECT id, temperature, humidity FROM temperature_and_humidity"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.TemperatureAndHumidity
	for rows.Next() {
		var sensor entities.TemperatureAndHumidity
		if err := rows.Scan(&sensor.ID, &sensor.Temperature, &sensor.Humidity); err != nil {
			return nil, fmt.Errorf("error al escanear los registros: %v", err)
		}
		records = append(records, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer los registros: %v", err)
	}

	return records, nil
}

// Delete elimina una medición por su ID
func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM temperature_and_humidity WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

// GetAverageTemperature calcula el promedio de la temperatura
func (m *MySQL) GetAverageTemperature() (float64, error) {
	var avgTemp float64
	query := "SELECT AVG(temperature) FROM temperature_and_humidity"
	err := m.conn.QueryRow(query).Scan(&avgTemp)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de temperatura: %v", err)
	}
	return avgTemp, nil
}

// GetAverageHumidity calcula el promedio de la humedad
func (m *MySQL) GetAverageHumidity() (float64, error) {
	var avgHum float64
	query := "SELECT AVG(humidity) FROM temperature_and_humidity"
	err := m.conn.QueryRow(query).Scan(&avgHum)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de humedad: %v", err)
	}
	return avgHum, nil
}
