package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/qwerty22121998/slink/pkg/models"
)

type ShortLinkRepository interface {
	Find(ctx context.Context, id string) (*models.ShortLink, error)
	Set(ctx context.Context, id string, link *models.ShortLink) error
}

type shortLinkRepository struct {
	col *firestore.CollectionRef
}

func NewShortLinkRepository(
	db *firestore.Client,
) ShortLinkRepository {
	return &shortLinkRepository{
		col: db.Collection("short"),
	}
}

func (r *shortLinkRepository) Find(ctx context.Context, id string) (*models.ShortLink, error) {
	var res *models.ShortLink
	doc, err := r.col.Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := doc.DataTo(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *shortLinkRepository) Set(ctx context.Context, id string, link *models.ShortLink) error {
	_, err := r.col.Doc(id).Set(ctx, link)
	if err != nil {
		return err
	}
	return nil
}
