package repository

import (
	"context"
	"inventory/internal/entity"
)

const (
	qryInsertUser = `
	INSERT INTO user (email, name, password) 
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
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, &u, qryGetUserByEmail, email) // ejecuta el query, mapea las columnas a los parametros o propiedades de un struct que estemos pasando
	if err!= nil {
        return nil, err
    }
	
	return u, err

}
