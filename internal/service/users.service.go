package service

import (
	"context"
	"errors"
	"inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists") // es buena practica no poner en mayuscula los errores ni ponerles punto
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email) // cuando vayamos a registrar un usuario revisamos que no exista en la base de datos
	if u != nil {
		return ErrUserAlreadyExists
	}

	//TODO: hash password

	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//TODO: decrypt password

	if u.Password != password {
		return nil, ErrInvalidCredentials

	}
	return &models.User{}, nil

}
