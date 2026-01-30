-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    role VARCHAR(15) NOT NULL,
    is_active BOOLEAN NOT NULL
);

INSERT INTO users (chat_id, name, role, is_active) VALUES
    (855343694, 'Михаил Пан-Асовский', 'admin', true);

ALTER TABLE bookings
    ADD COLUMN confirmed_by VARCHAR(50),
    ADD COLUMN cancelled_by VARCHAR(50),
    ADD COLUMN notes TEXT;


-- +goose Down
DROP TABLE IF EXISTS users;
ALTER TABLE bookings DROP COLUMN confirmed_by;
ALTER TABLE bookings DROP COLUMN cancelled_by;
ALTER TABLE bookings DROP COLUMN notes;

