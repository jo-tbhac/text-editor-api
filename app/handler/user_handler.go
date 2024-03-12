package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type IUserHandler interface {
	Create(c echo.Context) error
}

func Create(c echo.Context) error {
	return c.String(http.StatusCreated, "Created")
}
