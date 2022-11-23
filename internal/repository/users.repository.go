package repository

import (
	"context"
	"inventory/internal/entity"
)

const (
	qryInsertUser = `
	INSERT INTO USERS (email, name, password) 
	VALUES (?, ?, ?) 
`

	qryGetUserByEmail = `
	SELECT 
		id, 
		email, 
		name, 
		password
	FROM USERS
		where email = ?;
`

	qryInsertUserRole = `
	insert into USER_ROLES (user_id, role_id) values (:user_id, :role_id);`

	qryRemoveUserRole = `
	delete from USER_ROLES where user_id = :user_id and role_id= :role_id;`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {

	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, &u, qryGetUserByEmail, email) // ejecuta el query, mapea las columnas a los parametros o propiedades de un struct que estemos pasando
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)
	return err
}
