package shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"strings"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

func GenerateShortURL(urlOriginal dto.OriginalURL) (dto.ShortURL, error) {
	if urlOriginal == "" {
		return "", errors.New("original URL cannot be empty")
	}

	hash := sha256.Sum256([]byte(urlOriginal))

	encoded := base64.URLEncoding.EncodeToString(hash[:])

	encoded = strings.TrimRight(encoded, "=")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	encoded = strings.ReplaceAll(encoded, "+", "_")
	shortURL := encoded[:10]

	log.Printf("Generated short URL: %v", shortURL)
	return dto.ShortURL(shortURL), nil
}
