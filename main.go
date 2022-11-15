package main

import (
	"inventory/settings"

	"go.uber.org/fx"
	//"log"
)

// libreria para inyeccion de dependencias uber
func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
		), // pasamos todas la funciones que nos devuelvan un struct
		fx.Invoke(
		/* func (s *settings.Settings)  {
			log.Println(s)
		}, */
		), // ejecutar algun comando que necesitemos justo antes que la aplicacion empiece a correr
	)

	app.Run()

}
