package adapters

import (
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"
	"database/sql"
	"errors"
	"fmt"

	"Multidiciplinario/src/sensor_giroscopio/domain"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.IGyroscopeSensor {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.GyroscopeSensor) (entities.GyroscopeSensor, error) {
	query := `INSERT INTO gyroscope_sensor (x_axis, y_axis, z_axis) VALUES (?, ?, ?)`
	_, err := m.conn.Exec(query, sensor.XAxis, sensor.YAxis, sensor.ZAxis)
	if err != nil {
		return entities.GyroscopeSensor{}, err
	}
	return sensor, nil
}

func (m *MySQL) GetMeasurementByID(id int) (entities.GyroscopeSensor, error) {
	var sensor entities.GyroscopeSensor
	query := `SELECT id, x_axis, y_axis, z_axis FROM gyroscope_sensor WHERE id = ?`
	err := m.conn.QueryRow(query, id).Scan(&sensor.ID, &sensor.XAxis, &sensor.YAxis, &sensor.ZAxis)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetLatestMeasurement() (entities.GyroscopeSensor, error) {
	var sensor entities.GyroscopeSensor
	query := `SELECT id, x_axis, y_axis, z_axis FROM gyroscope_sensor ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query).Scan(&sensor.ID, &sensor.XAxis, &sensor.YAxis, &sensor.ZAxis)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAllMeasurements() ([]entities.GyroscopeSensor, error) {
	query := "SELECT id, x_axis, y_axis, z_axis FROM gyroscope_sensor"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.GyroscopeSensor
	for rows.Next() {
		var sensor entities.GyroscopeSensor
		if err := rows.Scan(&sensor.ID, &sensor.XAxis, &sensor.YAxis, &sensor.ZAxis); err != nil {
			return nil, fmt.Errorf("error al escanear los registros: %v", err)
		}
		records = append(records, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer los registros: %v", err)
	}

	return records, nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM gyroscope_sensor WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

func (m *MySQL) GetAverageGyroscopeData() (float64, error) {
	var avgGyroData float64
	query := "SELECT AVG(x_axis), AVG(y_axis), AVG(z_axis) FROM gyroscope_sensor"
	err := m.conn.QueryRow(query).Scan(&avgGyroData)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de los datos del giroscopio: %v", err)
	}
	return avgGyroData, nil
}
