CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    chatID BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    role VARCHAR(15) NOT NULL,
    is_active BOOLEAN NOT NULL
);

INSERT INTO users (chatID, name, role, is_active) VALUES
    (855343694, 'Михаил Пан-Асовский', 'admin', true);

ALTER TABLE bookings
    ADD COLUMN confirmed_by VARCHAR(50),
    ADD COLUMN cancelled_by VARCHAR(50),
    ADD COLUMN notes TEXT;