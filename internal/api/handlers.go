package api

import (
	"inventory/encryption"
	"inventory/internal/api/dtos"
	"inventory/internal/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type responseMenssage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMenssage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMenssage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}
		return c.JSON(http.StatusInternalServerError, responseMenssage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{} // parametros dto

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMenssage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMenssage{Message: err.Error()})
	}

	// buscar usuario en service
	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMenssage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMenssage{Message: "Internal server error"})
	}


	//TODO: return JWT token

	return c.JSON(http.StatusOK, map[string]string{"success": token})
}
