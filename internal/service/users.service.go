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
	ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	ErrRoleNotFound       = errors.New("role not found")
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

	/* if u.Password != password {
		return nil, ErrInvalidCredentials

	} */
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil

}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)

}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true 
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}        // si no encontramos el rol

	return s.repo.RemoveUserRole(ctx, userID, roleID)

}
