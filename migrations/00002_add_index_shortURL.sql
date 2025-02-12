-- +goose Up
-- Добавление индекса на поле shortURL
CREATE INDEX idx_shortURL ON url_db (shortURL);

-- +goose Down
-- Удаление индекса
DROP INDEX IF EXISTS idx_shortURL;