package adapters

import (
	"database/sql"
	"errors"
	"fmt"

	"Multidiciplinario/src/sensor_luz/domain"
	"Multidiciplinario/src/sensor_luz/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.ILightSensor {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.LightSensorLDR) (entities.LightSensorLDR, error) {
	query := `INSERT INTO light_sensor (lux, user_id) VALUES (?, ?)`
	_, err := m.conn.Exec(query, sensor.Lux, sensor.UserID)
	if err != nil {
		return entities.LightSensorLDR{}, err
	}
	return sensor, nil
}

func (m *MySQL) GetMeasurementByID(id int, userID int) (entities.LightSensorLDR, error) {
	var sensor entities.LightSensorLDR
	query := `SELECT id, lux, user_id FROM light_sensor WHERE id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(&sensor.ID, &sensor.Lux, &sensor.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetLatestMeasurement(userID int) (entities.LightSensorLDR, error) {
	var sensor entities.LightSensorLDR
	query := `SELECT id, lux, user_id FROM light_sensor WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query, userID).Scan(&sensor.ID, &sensor.Lux, &sensor.UserID)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAllMeasurements(userID int) ([]entities.LightSensorLDR, error) {
	query := "SELECT id, lux, user_id FROM light_sensor WHERE user_id = ?"
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.LightSensorLDR
	for rows.Next() {
		var sensor entities.LightSensorLDR
		if err := rows.Scan(&sensor.ID, &sensor.Lux, &sensor.UserID); err != nil {
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
	query := "DELETE FROM light_sensor WHERE id = ? AND user_id = ?"
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

func (m *MySQL) GetAverageLux(userID int) (float64, error) {
	var avgLux float64
	query := "SELECT AVG(lux) FROM light_sensor WHERE user_id = ?"
	err := m.conn.QueryRow(query, userID).Scan(&avgLux)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de lux: %v", err)
	}
	return avgLux, nil
}
