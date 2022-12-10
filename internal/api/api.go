package api

import (
	"inventory/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(serv service.Service) *API {
	return &API{
		serv:          serv,
		dataValidator: validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string) error { // esta función dejará de ejecutarse hasta que la aplicación termine o encuentre un error
	a.RegisterRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500"},
		AllowMethods: []string{echo.POST},
		AllowHeaders: []string{echo.HeaderContentType},
	}))

	return e.Start(address)
}
