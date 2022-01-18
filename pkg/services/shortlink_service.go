package services

import (
	"context"
	"github.com/qwerty22121998/slink/pkg/dtos"
	"github.com/qwerty22121998/slink/pkg/errors"
	"github.com/qwerty22121998/slink/pkg/models"
	"github.com/qwerty22121998/slink/pkg/repositories"
	"github.com/qwerty22121998/slink/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ShortLinkService interface {
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
	_, err := s.repo.Find(ctx, req.Short)
	if err == nil {
		return nil, errors.BadRequest("Short link %v already existed", req.Short)
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
	if link.Short == "" {
		link.Short = utils.Random(10)
	}
	if req.ExpiredAt != 0 {
		expiredAt := time.Unix(req.ExpiredAt, 0)
		link.ExpiredAt = &expiredAt
	}

	if err := s.repo.Set(ctx, link.Short, link); err != nil {
		return nil, err
	}
	return &dtos.LinkShortenerResponse{
		Message: "",
		Data:    nil,
	}

}
