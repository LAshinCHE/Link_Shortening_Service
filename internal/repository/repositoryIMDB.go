package repository

import (
	"errors"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/repository/schema"
)

func NewRepositoryIMDB() *InMemoryURLStorage {
	return &InMemoryURLStorage{
		urlStorage: make(map[schema.UrlShortSchema]schema.OriginalURLSchema),
	}
}

type InMemoryURLStorage struct {
	urlStorage map[schema.UrlShortSchema]schema.OriginalURLSchema
}

// TODO - реализвовать логику добавления URL в базу
func (imdb *InMemoryURLStorage) AddURL(urlOriginalDto dto.OriginalURL, urlShortDto dto.ShortURL) error {

	shortURL := toSChemaShort(urlShortDto)
	originalURL := toSChemaOriginal(urlOriginalDto)

	if _, ok := imdb.urlStorage[shortURL]; ok {
		return ErrorExistURL
	}
	imdb.urlStorage[shortURL] = originalURL
	return nil
}

// TODO - реализвовать логику взятия URL из базу
func (imdb *InMemoryURLStorage) GetURL(urlShortDto dto.ShortURL) (dto.OriginalURL, error) {

	shortURL := toSChemaShort(urlShortDto)

	originalURL, ok := imdb.urlStorage[shortURL]
	if !ok {
		return "", ErorrNoSuchURLInIMDB
	}

	urlOriginalDto := toOriginalDTO(originalURL)
	return urlOriginalDto, nil
}

func toSChemaShort(url dto.ShortURL) schema.UrlShortSchema {
	return schema.UrlShortSchema(url)
}

func toOriginalDTO(url schema.OriginalURLSchema) dto.OriginalURL {
	return dto.OriginalURL(url)
}

func toSChemaOriginal(url dto.OriginalURL) schema.OriginalURLSchema {
	return schema.OriginalURLSchema(url)
}

var (
	ErrorExistURL        = errors.New("This error is already exist")
	ErorrNoSuchURLInIMDB = errors.New("No such url in Imdb")
)
