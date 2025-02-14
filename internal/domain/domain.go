package domain

import "github.com/LAshinCHE/Link_Shortening_Service/internal/dto"

type Repository interface {
	AddURL(urlOriginal dto.OriginalURL, urlShort dto.ShortURL) (error, *dto.ShortURL)
	GetURL(urlShort dto.ShortURL) (error, *dto.OriginalURL)
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
