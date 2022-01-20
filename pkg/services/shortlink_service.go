package services

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/slink/pkg/dtos"
	"github.com/qwerty22121998/slink/pkg/models"
	"github.com/qwerty22121998/slink/pkg/repositories"
	"github.com/qwerty22121998/slink/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type ShortLinkService interface {
	Create(ctx context.Context, req dtos.ShortLink) (*dtos.LinkShortenerResponse, error)
	Find(ctx context.Context, short string) (*dtos.ShortLink, error)
}

type shortLinkService struct {
	repo repositories.ShortLinkRepository
}

func NewShortLinkService(
	repo repositories.ShortLinkRepository,
) ShortLinkService {
	return &shortLinkService{
		repo: repo,
	}
}

func (s *shortLinkService) Create(ctx context.Context, req dtos.ShortLink) (*dtos.LinkShortenerResponse, error) {
	if req.Short == "" {
		req.Short = utils.Random(10)
	}
	_, err := s.repo.Find(ctx, req.Short)
	if err == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Short link %v already existed", req.Short))
	}
	if err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	link := &models.ShortLink{
		CreatedAt:  time.Now(),
		Name:       req.Name,
		URL:        req.URL,
		Short:      req.Short,
		ClickCount: 0,
	}
	if link.Name == "" {
		link.Name = link.URL
	}
	if req.ExpiredAt != 0 {
		expiredAt := time.Unix(req.ExpiredAt, 0)
		link.ExpiredAt = &expiredAt
	}

	if err := s.repo.Set(ctx, link.Short, link); err != nil {
		return nil, err
	}
	return &dtos.LinkShortenerResponse{
		Message: "success",
		Data: &dtos.ShortLink{
			Name:      link.Name,
			URL:       link.URL,
			Short:     link.Short,
			ExpiredAt: req.ExpiredAt,
		},
	}, nil
}

func (s *shortLinkService) Find(ctx context.Context, short string) (*dtos.ShortLink, error) {
	link, err := s.repo.Find(ctx, short)

	if err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, err
	}

	expiredAt := int64(0)
	if link.ExpiredAt != nil {
		expiredAt = link.ExpiredAt.Unix()
	}
	return &dtos.ShortLink{
		Name:      link.Name,
		URL:       link.URL,
		Short:     link.Short,
		ExpiredAt: expiredAt,
	}, nil
}
