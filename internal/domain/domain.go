package domain

import (
	"context"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

type Repository interface {
	AddURL(urlOriginal dto.OriginalURL, urlShort dto.ShortURL) (error, *dto.ShortURL)
	GetURL(urlShort dto.ShortURL) (error, *dto.OriginalURL)
}

type ShortenerService struct {
	repo Repository
}

func NewShortenerService(r Repository) *ShortenerService {
	return &ShortenerService{
		repo: r,
	}
}

func AddURL(ctx context.Context, url dto.OriginalURL) (*dto.ShortURL, error) {
	return nil, nil
}
func GetURL(ctx context.Context, url dto.ShortURL) (*dto.OriginalURL, error) {
	return nil, nil
}
