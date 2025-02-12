package storage

import "github.com/LAshinCHE/Link_Shortening_Service/internal/dto"

func NewPostgraceURLDB() *PostgraceURLDB {
	return &PostgraceURLDB{}
}

type PostgraceURLDB struct {
}

// TODO - реализвовать логику добавления URL в базу
func (mus *PostgraceURLDB) AddURL(url dto.OriginalURL) (error, *dto.ShortURL) {
	return nil, nil
}

// TODO - реализвовать логику взятия URL из базу
func (mus *PostgraceURLDB) GetURL(url dto.ShortURL) (error, *dto.OriginalURL) {
	return nil, nil
}
