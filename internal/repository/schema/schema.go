package schema

type URL struct {
	ID          int64  `db:"id"`
	ShortURL    string `db:"shortURL"`
	OriginalURL string `db:"originalURL"`
}

type (
	UrlShortSchema    string
	OriginalURLSchema string
)
