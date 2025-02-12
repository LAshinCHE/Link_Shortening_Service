package storage

import "github.com/LAshinCHE/Link_Shortening_Service/internal/dto"

type URLStorage interface {
	AddURL(url dto.OriginalURL) *dto.ShortURL
	GetURL(url dto.ShortURL) *dto.OriginalURL
}
