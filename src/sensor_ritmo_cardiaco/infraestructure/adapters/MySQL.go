package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.IHeartRate {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.HeartRate) (entities.HeartRate, error) {
	query := `INSERT INTO heart_rate (heart_rate, user_id) VALUES (?, ?)`
	_, err := m.conn.Exec(query, sensor.HeartRate, sensor.UserID)
	if err != nil {
		return entities.HeartRate{}, err
	}
	return sensor, nil
}

func (m *MySQL) GetMeasurementByID(id int, userID int) (entities.HeartRate, error) {
	var sensor entities.HeartRate
	query := `SELECT id, heart_rate, user_id FROM heart_rate WHERE id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(&sensor.ID, &sensor.HeartRate, &sensor.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetLatestMeasurement(userID int) (entities.HeartRate, error) {
	var sensor entities.HeartRate
	query := `SELECT id, heart_rate, user_id FROM heart_rate WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query, userID).Scan(&sensor.ID, &sensor.HeartRate, &sensor.UserID)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAllMeasurements(userID int) ([]entities.HeartRate, error) {
	query := "SELECT id, heart_rate, user_id FROM heart_rate WHERE user_id = ?"
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.HeartRate
	for rows.Next() {
		var sensor entities.HeartRate
		if err := rows.Scan(&sensor.ID, &sensor.HeartRate, &sensor.UserID); err != nil {
			return nil, fmt.Errorf("error al escanear los registros: %v", err)
		}
		records = append(records, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer los registros: %v", err)
	}

	return records, nil
}

func (m *MySQL) Delete(id int, userID int) error {
	query := "DELETE FROM heart_rate WHERE id = ? AND user_id = ?"
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

func (m *MySQL) GetAverageHeartRate(userID int) (float64, error) {
	var avgHeartRate float64
	query := "SELECT AVG(heart_rate) FROM heart_rate WHERE user_id = ?"
	err := m.conn.QueryRow(query, userID).Scan(&avgHeartRate)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de ritmo cardiaco: %v", err)
	}
	return avgHeartRate, nil
}
