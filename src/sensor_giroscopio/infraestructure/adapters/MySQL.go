package adapters

import (
	"Multidiciplinario/src/sensor_giroscopio/domain"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.IGyroscopeSensor {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.GyroscopeSensor) (entities.GyroscopeSensor, error) {
	query := `INSERT INTO gyroscope (x_axis, y_axis, z_axis, user_id) VALUES (?, ?, ?, ?)`
	log.Printf("Guardando sensor: X=%.2f, Y=%.2f, Z=%.2f, UserID=%d", sensor.XAxis, sensor.YAxis, sensor.ZAxis, sensor.UserID)
	_, err := m.conn.Exec(query, sensor.XAxis, sensor.YAxis, sensor.ZAxis, sensor.UserID)
	if err != nil {
		log.Printf("Error al guardar el sensor: %v", err)
		return entities.GyroscopeSensor{}, err
	}
	log.Println("Sensor guardado exitosamente")
	return sensor, nil
}

func (m *MySQL) GetMeasurementByID(id int, userID int) (entities.GyroscopeSensor, error) {
	var sensor entities.GyroscopeSensor
	query := `SELECT id, x_axis, y_axis, z_axis, user_id FROM gyroscope WHERE id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(&sensor.ID, &sensor.XAxis, &sensor.YAxis, &sensor.ZAxis, &sensor.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetLatestMeasurement(userID int) (entities.GyroscopeSensor, error) {
	var sensor entities.GyroscopeSensor
	query := `SELECT id, x_axis, y_axis, z_axis, user_id FROM gyroscope WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query, userID).Scan(&sensor.ID, &sensor.XAxis, &sensor.YAxis, &sensor.ZAxis, &sensor.UserID)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAllMeasurements(userID int) ([]entities.GyroscopeSensor, error) {
	query := "SELECT id, x_axis, y_axis, z_axis, user_id FROM gyroscope WHERE user_id = ?"
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.GyroscopeSensor
	for rows.Next() {
		var sensor entities.GyroscopeSensor
		if err := rows.Scan(&sensor.ID, &sensor.XAxis, &sensor.YAxis, &sensor.ZAxis, &sensor.UserID); err != nil {
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
	query := "DELETE FROM gyroscope WHERE id = ? AND user_id = ?"
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

func (m *MySQL) GetAverageGyroscopeData(userID int) (float64, error) {
	var avgX, avgY, avgZ float64
	query := `SELECT AVG(x_axis), AVG(y_axis), AVG(z_axis) FROM gyroscope WHERE user_id = ?`
	err := m.conn.QueryRow(query, userID).Scan(&avgX, &avgY, &avgZ)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de los datos del giroscopio: %v", err)
	}

	avgGyroData := (avgX + avgY + avgZ) / 3
	return avgGyroData, nil
}
