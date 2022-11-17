package database

import (
	"context"
	"fmt"
	"inventory/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true", // puerto es un numero asi que se pone %d
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	) // el parseTime=True se usa para cuando se va a trabajar con las fechas en golang

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
