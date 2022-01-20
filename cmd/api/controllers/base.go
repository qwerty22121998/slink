package controllers

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller interface {
	Register(c *echo.Group)
}

type controller struct {
	validator *validator.Validate
}

func newController() controller {
	return controller{
		validator: validator.New(),
	}
}

func (c *controller) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
