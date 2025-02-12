-- +goose Up
-- Создание таблицы url_db
CREATE TABLE url_db (
    id SERIAL PRIMARY KEY,
    shortURL TEXT NOT NULL UNIQUE,
    originalURL TEXT NOT NULL
);

-- +goose Down
-- Удаление таблицы url_db
DROP TABLE IF EXISTS url_db;