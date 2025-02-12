package storage

import (
	"sync"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

func NewInMemoryURLStorage() *InMemoryURLStorage {
	return &InMemoryURLStorage{
		urlStorage: make(map[dto.ShortURL]*dto.OriginalURL),
	}
}

type InMemoryURLStorage struct {
	lock       sync.RWMutex
	urlStorage map[dto.ShortURL]*dto.OriginalURL
}

// TODO - реализвовать логику добавления URL в базу
func (mus *InMemoryURLStorage) AddURL(url dto.OriginalURL) (error, *dto.ShortURL) {
	return nil, nil
}

// TODO - реализвовать логику взятия URL из базу
func (mus *InMemoryURLStorage) GetURL(url dto.ShortURL) (error, *dto.OriginalURL) {
	return nil, nil
}
