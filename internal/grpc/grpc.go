package grpc

import (
	"context"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

type shortener interface {
	AddURL(url dto.OriginalURL) (*dto.ShortURL, error)
	GetURL(url dto.ShortURL) (*dto.OriginalURL, error)
}

type Handler struct {
	shortener shortener
}

// TODO - сделать метод Run для grpc сервиса, который будет запускать наш сервис
func Run(ctx context.Context, app shortener) {

}
