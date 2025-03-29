package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"Multidiciplinario/src/sensor_temperatura_corporal/domain"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.IBodyTemperature {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.BodyTemperature) (entities.BodyTemperature, error) {
	query := `INSERT INTO body_temperature (temperature, user_id) VALUES (?, ?)`
	_, err := m.conn.Exec(query, sensor.Temperature, sensor.UserID)
	if err != nil {
		return entities.BodyTemperature{}, err
	}
	return sensor, nil
}

func (m *MySQL) GetMeasurementByID(id, userID int) (entities.BodyTemperature, error) {
	var sensor entities.BodyTemperature
	query := `SELECT id, temperature, user_id FROM body_temperature WHERE id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(&sensor.ID, &sensor.Temperature, &sensor.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetLatestMeasurement(userID int) (entities.BodyTemperature, error) {
	var sensor entities.BodyTemperature
	query := `SELECT id, temperature, user_id FROM body_temperature WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query, userID).Scan(&sensor.ID, &sensor.Temperature, &sensor.UserID)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAllMeasurements(userID int) ([]entities.BodyTemperature, error) {
	query := "SELECT id, temperature, user_id FROM body_temperature WHERE user_id = ?"
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.BodyTemperature
	for rows.Next() {
		var sensor entities.BodyTemperature
		if err := rows.Scan(&sensor.ID, &sensor.Temperature, &sensor.UserID); err != nil {
			return nil, fmt.Errorf("error al escanear los registros: %v", err)
		}
		records = append(records, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer los registros: %v", err)
	}

	return records, nil
}

func (m *MySQL) Delete(id, userID int) error {
	query := "DELETE FROM body_temperature WHERE id = ? AND user_id = ?"
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

func (m *MySQL) GetAverageTemperature(userID int) (float64, error) {
	var avgTemp float64
	query := "SELECT AVG(temperature) FROM body_temperature WHERE user_id = ?"
	err := m.conn.QueryRow(query, userID).Scan(&avgTemp)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de temperatura: %v", err)
	}
	return avgTemp, nil
}
