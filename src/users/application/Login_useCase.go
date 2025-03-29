package application

import (
	"Multidiciplinario/src/core/security"
	"Multidiciplinario/src/users/domain"
	"Multidiciplinario/src/users/domain/entities"
	"errors"
)

type AuthService struct {
	clientRepo domain.IUser
}

func NewAuthService(clientRepo domain.IUser) *AuthService {
	return &AuthService{clientRepo: clientRepo}
}

func (as *AuthService) Login(email, password string) (map[string]interface{}, error) {
	client, err := as.clientRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	if !security.CheckPassword(client.Password, password) {
		return nil, errors.New("contraseña incorrecta")
	}
	token, err := security.GenerateJWT(int(client.ID), client.Email)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":  token,
		"userId": client.ID,
	}, nil
}

func (as *AuthService) Register(client entities.User) error {
	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword
	return as.clientRepo.Save(client)
}
