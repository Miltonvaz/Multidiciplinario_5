package application

import (
	"Multidiciplinario/src/core/security"
	"Multidiciplinario/src/users/application/repositories"
	"Multidiciplinario/src/users/domain"
	"Multidiciplinario/src/users/domain/entities"
	"fmt"
	"log"
)

type CreateUser struct {
	db                  domain.IUser
	serviceNotification *repositories.ServiceNotification
}

func NewCreateUser(db domain.IUser, serviceNotification *repositories.ServiceNotification) *CreateUser {
	return &CreateUser{
		db:                  db,
		serviceNotification: serviceNotification,
	}
}

func (cc *CreateUser) Execute(client entities.User) error {
	existingUser, err := cc.db.GetByEsp32Serial(client.Id_esp32)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("el número de serie del ESP32 ya está en uso, por favor ingrese otro")
	}

	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword

	if err := cc.db.Save(client); err != nil {
		return err
	}

	err = cc.serviceNotification.NotifyAppointmentCreated(client)
	if err != nil {
		log.Printf("Error notificando la creación del usuario: %v", err)
	}

	return nil
}
