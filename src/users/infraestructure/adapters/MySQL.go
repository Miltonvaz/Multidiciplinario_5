package adapters

import (
	"Multidiciplinario/src/users/domain/entities"
	"database/sql"
	"errors"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(user entities.User) error {
	// Verificar si el número de serie ya está registrado en esp32_devices
	existingDevice, err := m.GetByEsp32Serial(user.Id_esp32)
	if err != nil && err.Error() != "device not found" {
		return fmt.Errorf("error checking ESP32 serial: %v", err)
	}
	if existingDevice != nil {
		return errors.New("ESP32 serial number already in use. Please enter a different one.")
	}

	// Insertar el número de serie en la tabla esp32_devices si no existe
	err = m.InsertEsp32Serial(user.Id_esp32)
	if err != nil {
		return fmt.Errorf("error inserting ESP32 serial into the devices table: %v", err)
	}

	// Ahora guardamos al usuario en la tabla de usuarios
	query := `INSERT INTO users (name, last_name, email, backup_email, age, password, esp32_serial) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = m.conn.Exec(query, user.Name, user.LastName, user.Email, user.BackupEmail, user.Age, user.Password, user.Id_esp32)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	return nil
}

func (m *MySQL) InsertEsp32Serial(serial string) error {
	// Asegurarse de que el serial se inserte en la tabla esp32_devices
	query := `INSERT INTO esp32_devices (serial_number) VALUES (?)`
	_, err := m.conn.Exec(query, serial)
	if err != nil {
		return fmt.Errorf("failed to insert ESP32 serial: %v", err)
	}
	return nil
}

func (m *MySQL) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	query := `SELECT id, name, last_name, email, backup_email, age, password, esp32_serial FROM users WHERE email = ? LIMIT 1`
	err := m.conn.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}

	return user, nil
}

func (m *MySQL) GetByEsp32Serial(serial string) (*entities.User, error) {
	var user entities.User
	// Asegurarse de que se esté buscando correctamente por el serial
	query := `SELECT id, name, last_name, email, backup_email, age, password, esp32_serial FROM users WHERE esp32_serial = ? LIMIT 1`
	err := m.conn.QueryRow(query, serial).Scan(
		&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No se encontró usuario con ese número de serie
		}
		return nil, fmt.Errorf("failed to check ESP32 serial: %v", err)
	}

	return &user, nil
}

func (m *MySQL) GetAll() ([]entities.User, error) {
	query := "SELECT id, name, last_name, email, backup_email, age, password, esp32_serial FROM users"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

func (m *MySQL) GetById(id int) (entities.User, error) {
	query := "SELECT id, name, last_name, email, backup_email, age, password, esp32_serial FROM users WHERE id = ?"
	row := m.conn.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.BackupEmail, &user.Age, &user.Password, &user.Id_esp32)
	if err == sql.ErrNoRows {
		return entities.User{}, errors.New("user not found")
	} else if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return user, nil
}

func (m *MySQL) Edit(user entities.User) error {
	query := "UPDATE users SET name = ?, last_name = ?, email = ?, backup_email = ?, age = ?, password = ? WHERE id = ?"
	_, err := m.conn.Exec(query, user.Name, user.LastName, user.Email, user.BackupEmail, user.Age, user.Password, user.ID)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
