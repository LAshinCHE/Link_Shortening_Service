package domain

import "github.com/LAshinCHE/Link_Shortening_Service/internal/dto"

type Repository interface {
	AddURL(url dto.OriginalURL) (error, *dto.ShortURL)
	GetURL(url dto.ShortURL) (error, *dto.OriginalURL)
}

type Deps struct {
	Repository
}

type ShortenerService struct {
	Deps
}

func NewShortenerService(d Deps) *ShortenerService {
	return &ShortenerService{
		d,
	}
}
