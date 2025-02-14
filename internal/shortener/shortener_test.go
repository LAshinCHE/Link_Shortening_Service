package shortener

import (
	"strings"
	"testing"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
)

func TestGenerateShortURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		originalURL dto.OriginalURL
		wantErr     bool
	}{
		{
			name:        "Valid URL",
			originalURL: "https://www.example.com",
			wantErr:     false,
		},
		{
			name:        "Empty URL",
			originalURL: "",
			wantErr:     true,
		},
		{
			name:        "Long URL",
			originalURL: "https://www.example.com/very/long/url/with/many/parts",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			shortURL, err := GenerateShortURL(tt.originalURL)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(shortURL) != 10 {
					t.Errorf("Expected short URL length 10, got %d", len(shortURL))
				}

				validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
				for _, char := range shortURL {
					if !strings.ContainsRune(validChars, char) {
						t.Errorf("Invalid character in short URL: %c", char)
					}
				}
			}
		})
	}
}

func TestGenerateShortURL_Uniqueness(t *testing.T) {
	url1 := "https://www.example.com/first"
	url2 := "https://www.example.com/second"

	shortURL1, err1 := GenerateShortURL(dto.OriginalURL(url1))
	shortURL2, err2 := GenerateShortURL(dto.OriginalURL(url2))

	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error: %v, %v", err1, err2)
	}

	if shortURL1 == shortURL2 {
		t.Errorf("Expected unique short URLs, got the same: %s", shortURL1)
	}
}

func TestGenerateShortURL_Consistency(t *testing.T) {
	url := "https://www.example.com/consistent"

	shortURL1, err1 := GenerateShortURL(dto.OriginalURL(url))
	shortURL2, err2 := GenerateShortURL(dto.OriginalURL(url))

	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error: %v, %v", err1, err2)
	}

	if shortURL1 != shortURL2 {
		t.Errorf("Expected consistent short URLs, got different: %s and %s", shortURL1, shortURL2)
	}
}
