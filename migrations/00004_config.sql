-- +goose Up
CREATE TABLE config (
    auto_confirm BOOLEAN NOT NULL
);

INSERT INTO config (auto_confirm) VALUES (true);

-- +goose Down
DROP TABLE IF EXISTS config;