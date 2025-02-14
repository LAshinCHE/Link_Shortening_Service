package repository

import "github.com/LAshinCHE/Link_Shortening_Service/internal/dto"

type URLStorage interface {
	AddURL(urlOriginal dto.OriginalURL, urlShort dto.ShortURL) error
	GetURL(url dto.ShortURL) (*dto.OriginalURL, error)
}
