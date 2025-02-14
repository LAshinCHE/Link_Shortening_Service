package repository

import (
	"context"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/dto"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/repository/schema"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/shortener"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

const (
	URLDB = "url_db"
)

var (
	urlColumns = []string{"id", "shortURL", "originalURL"}
)

type RepositoryPg struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewRepositoryPg(ctx context.Context, connString string) (*RepositoryPg, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &RepositoryPg{ctx: ctx, db: conn}, nil
}

func (r *RepositoryPg) GetURL(urlShort dto.ShortURL) (*dto.OriginalURL, error) {
	query := sq.Select(urlColumns...).
		From(URLDB).
		Where("shortURL = $1", urlShort).PlaceholderFormat(sq.Dollar)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var url schema.URL
	if err := pgxscan.Select(r.ctx, r.db, url, rawQuery, args); err != nil {
		return nil, err
	}
	originalURL := urlToOriginal(url)
	return &originalURL, nil
}

func (r *RepositoryPg) AddURL(urlOriginal dto.OriginalURL, urlShort dto.ShortURL) (*dto.ShortURL, error) {
	shortURL, err := shortener.GenerateShortURL(urlOriginal)

	if err != nil {
		return nil, err
	}

	query := sq.Insert(URLDB).
		Columns("shortURL", "originalURL").
		Values(*shortURL, urlOriginal).
		PlaceholderFormat(sq.Dollar)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(r.ctx, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	return shortURL, nil
}

func urlToOriginal(url schema.URL) dto.OriginalURL {
	return dto.OriginalURL(url.OriginalURL)
}

func urlToShort(url schema.URL) dto.ShortURL {
	return dto.ShortURL(url.ShortURL)
}
