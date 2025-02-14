package domain

import (
	"context"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/shortener"
)

type Repository interface {
	AddURL(urlOriginal dto.OriginalURL, urlShort dto.ShortURL) error
	GetURL(url dto.ShortURL) (dto.OriginalURL, error)
}

type ShortenerService struct {
	repo Repository
}

func NewShortenerService(r Repository) *ShortenerService {
	return &ShortenerService{
		repo: r,
	}
}

func (serv *ShortenerService) AddURL(ctx context.Context, url dto.OriginalURL) (dto.ShortURL, error) {
	shortUrl, err := shortener.GenerateShortURL(url)
	if err != nil {
		return "", err
	}
	err = serv.repo.AddURL(url, shortUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}
func (serv *ShortenerService) GetURL(ctx context.Context, url dto.ShortURL) (dto.OriginalURL, error) {
	return serv.repo.GetURL(url)
}
