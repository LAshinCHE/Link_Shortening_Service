package repository

import (
	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

func NewRepositoryIMDB() *InMemoryURLStorage {
	return &InMemoryURLStorage{
		urlStorage: make(map[dto.ShortURL]*dto.OriginalURL),
	}
}

type InMemoryURLStorage struct {
	urlStorage map[dto.ShortURL]*dto.OriginalURL
}

// TODO - реализвовать логику добавления URL в базу
func (imdb *InMemoryURLStorage) AddURL(url dto.OriginalURL) (*dto.ShortURL, error) {
	return nil, nil
}

// TODO - реализвовать логику взятия URL из базу
func (imdb *InMemoryURLStorage) GetURL(url dto.ShortURL) (*dto.OriginalURL, error) {
	return nil, nil
}
