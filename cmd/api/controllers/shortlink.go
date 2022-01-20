package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/slink/pkg/dtos"
	"github.com/qwerty22121998/slink/pkg/services"
	"net/http"
)

type ShortLinkController interface {
	Controller
}

type shortLinkController struct {
	controller
	service services.ShortLinkService
}

func NewShortLinkController(
	service services.ShortLinkService,
) ShortLinkController {
	return &shortLinkController{
		controller: newController(),
		service:    service,
	}
}

func (c *shortLinkController) Register(r *echo.Group) {
	r.POST("/short", c.LinkShorten)
	r.GET("/g/:short", c.Go)
}

func (c *shortLinkController) LinkShorten(ctx echo.Context) error {
	var req dtos.ShortLink
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	resp, err := c.service.Create(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (c *shortLinkController) Go(ctx echo.Context) error {
	short := ctx.Param("short")
	link, err := c.service.Find(ctx.Request().Context(), short)
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, link.URL)
}
