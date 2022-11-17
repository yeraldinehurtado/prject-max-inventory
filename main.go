package main

import (
	"context"
	"inventory/database"
	"inventory/internal/repository"
	"inventory/internal/service"
	"inventory/settings"

	"github.com/jmoiron/sqlx"
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
		), // pasamos todas la funciones que nos devuelvan un struct
		fx.Invoke(
			func(db *sqlx.DB){
				_, err := db.Query("SELECT * FROM USERS")
				if err!= nil {
                    panic(err)
				}
			},// revisar que la conexion haya sido exitosa
		/* func (s *settings.Settings)  {
			log.Println(s)
		}, */
		), // ejecutar algun comando que necesitemos justo antes que la aplicacion empiece a correr
	)

	app.Run()

}
