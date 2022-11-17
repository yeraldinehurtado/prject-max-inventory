package service

import (
	"context"
	"errors"
	"inventory/encryption"
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

	bb, err := encryption.Encrypt([]byte(password)) //arreglo de bytes de password y esto me retorna los bytes o un error
	if err != nil {
		return err

	}

	pass := encryption.ToBase64(bb) // luego lo tenemos que pasar a string
	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//TODO: decrypt password
	// descifrar la contraseña que viene desde la base de datos

	bb, err := encryption.FromBase64(u.Password) //
	if err != nil {
		return nil, err

	}

	decryptedPassword, err := encryption.Decrypt(bb) // descifrar la contr
	if err != nil {
		return nil, err

	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	if u.Password != password {
		return nil, ErrInvalidCredentials

	}
	return &models.User{
		ID:    int(u.ID),
		Email: u.Email,
		Name:  u.Name,
	}, nil

}
