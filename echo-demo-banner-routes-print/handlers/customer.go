package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerHandler interface {
	GetById(echo echo.Context) error
}

func NewCustomerHandler() CustomerHandler {
	return &customerHandler{}
}

type customerHandler struct {
}

func (a customerHandler) GetById(e echo.Context) error {
	return e.JSON(http.StatusOK, "Customer Handler")
}
