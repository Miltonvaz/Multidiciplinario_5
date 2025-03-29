package application

import (
	"Multidiciplinario/src/core/security"
	"Multidiciplinario/src/users/domain"
	"Multidiciplinario/src/users/domain/entities"
)

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cc *CreateUser) Execute(client entities.User) error {
	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword
	return cc.db.Save(client)
}
