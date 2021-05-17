package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthHandler interface {
	Get(echo echo.Context) error
}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

type healthHandler struct {
}

func (a healthHandler) Get(e echo.Context) error {
	return e.JSON(http.StatusOK, "health Handler")
}
