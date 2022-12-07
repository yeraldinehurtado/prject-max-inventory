package api

import (
	"inventory/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
	return e.Start(address)
}
