package handels

import (
	"context"
	"errors"
	"log"
	"regexp"

	pb "github.com/LAshinCHE/Link_Shortening_Service/api/proto/shortener"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

var (
	ErrorZeroSizeURL          = errors.New("Url has zero len")
	ErrorWrongShortSizeURL    = errors.New("Short url must 10 symbols len")
	ErrorWrongShortURLSymbols = errors.New(`The URL must consist of lowercase
	and uppercase Latin characters, numbers, and the _ symbol (underscore).`)
)

type shortener interface {
	AddURL(ctx context.Context, url dto.OriginalURL) (dto.ShortURL, error)
	GetURL(ctx context.Context, url dto.ShortURL) (dto.OriginalURL, error)
}

type Handler struct {
	pb.UnimplementedShortenerServer
	service shortener
}

func NewHandler(serv shortener) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) AddURL(ctx context.Context, in *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	log.Printf("Received: %v", in.GetOriginalURL())
	if len(in.GetOriginalURL()) == 0 {
		return nil, ErrorZeroSizeURL
	}

	originalDTOURL := ToOriginalURLDTO(in.GetOriginalURL())
	shortDTOURL, err := h.service.AddURL(ctx, originalDTOURL)

	if err != nil {
		return nil, err
	}

	return &pb.AddURLResponse{ShortURL: string(shortDTOURL)}, nil
}

func (h *Handler) GetURL(ctx context.Context, in *pb.GetURLRequest) (*pb.GetURLResponse, error) {

	if len(in.GetShortURL()) < 10 {
		return nil, ErrorWrongShortSizeURL
	}

	pattern := `^[a-zA-Z0-9_]+$`
	matched, err := regexp.MatchString(pattern, in.GetShortURL())
	if err != nil || matched == false {
		return nil, ErrorWrongShortURLSymbols
	}

	shortDTOURL := ToShortURLDTO(in.GetShortURL())
	originalDTOURL, err := h.service.GetURL(ctx, shortDTOURL)

	if err != nil {
		return nil, err
	}

	return &pb.GetURLResponse{OriginalURL: string(originalDTOURL)}, nil
}

func ToOriginalURLDTO(url string) dto.OriginalURL {
	return dto.OriginalURL(url)
}

func ToShortURLDTO(url string) dto.ShortURL {
	return dto.ShortURL(url)
}
