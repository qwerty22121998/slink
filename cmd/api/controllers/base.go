package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	Register(c *echo.Group)
}
