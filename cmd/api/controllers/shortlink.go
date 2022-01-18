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
	service services.ShortLinkService
}

func NewShortLinkController(
	service services.ShortLinkService,
) ShortLinkController {
	return &shortLinkController{
		service: service,
	}
}

func (c *shortLinkController) Register(r *echo.Group) {
	r.GET("/short", c.LinkShorten)
}

func (c *shortLinkController) LinkShorten(ctx echo.Context) error {
	var req dtos.ShortLink
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.LinkShortenerResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dtos.LinkShortenerResponse{
		Message: "success",
	})
}
