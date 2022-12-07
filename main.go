package main

import (
	"context"
	"fmt"
	"inventory/database"
	"inventory/internal/api"
	"inventory/internal/repository"
	"inventory/internal/service"
	"inventory/settings"

	//"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	//"log"
)

// libreria para inyeccion de dependencias uber
func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New, // dependencia a la base de datos
			repository.New,
			service.New,
			api.New,
			echo.New,
		), // pasamos todas la funciones que nos devuelvan un struct
		fx.Invoke(
			setLifeCycle,
		/* func(db *sqlx.DB) {
			_, err := db.Query("SELECT * FROM USERS")
			if err != nil {
				panic(err)
			}
		}, */ // revisar que la conexion haya sido exitosa
		/* func (s *settings.Settings)  {
			log.Println(s)
		}, */
		/* func(ctx context.Context, serv service.Service) {
			err := serv.RegisterUser(ctx, "my@mail.com", "myname", "mypassword")
			if err != nil {
				panic(err)
			}

			u, err := serv.LoginUser(ctx, "my@mail.com", "mypassword")
			if err != nil {
				panic(err)
			}

			if u.Name != "myname" {
				panic("wrong name")

			}
		}, */ // funcion para revisar si mi codigo esta bien
		),    // ejecutar algun comando que necesitemos justo antes que la aplicacion empiece a correr

	)

	app.Run()

}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
