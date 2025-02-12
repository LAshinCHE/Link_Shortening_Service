package grpc

import (
	"context"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

type shortener interface {
	AddURL(url dto.OriginalURL) (error, *dto.ShortURL)
	GetURL(url dto.ShortURL) (error, *dto.OriginalURL)
}

type urlStorage interface {
	AddURL(url dto.OriginalURL) *dto.ShortURL
	GetURL(url dto.ShortURL) *dto.OriginalURL
}

type Handler struct {
	service shortener
	storage urlStorage
}

func Run(ctx context.Context, addr string, app shortener)
