package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type AccountHandler interface {
	List(echo echo.Context) error
}

func NewAccountHandler() AccountHandler {
	return &accountHandler{}
}

type accountHandler struct {
}

func (a accountHandler) List(e echo.Context) error {
	return e.JSON(http.StatusOK, "Account Handler")
}
