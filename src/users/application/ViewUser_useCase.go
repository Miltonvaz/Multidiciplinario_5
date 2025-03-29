package application

import (
	"Multidiciplinario/src/users/domain"
	"Multidiciplinario/src/users/domain/entities"
)

type ViewUser struct {
	db domain.IUser
}

func NewListUser(db domain.IUser) *ViewUser {
	return &ViewUser{db: db}
}

func (vc *ViewUser) Execute() ([]entities.User, error) {
	return vc.db.GetAll()
}
