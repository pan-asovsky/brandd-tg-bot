-- +goose Up
CREATE TABLE service_types (
    id BIGSERIAL PRIMARY KEY,
    service_code VARCHAR(20) NOT NULL UNIQUE,
    service_name VARCHAR(50) NOT NULL,
    is_composite BOOLEAN NOT NULL
);

CREATE TABLE prices (
    id BIGSERIAL PRIMARY KEY,
    rim_size INTEGER NOT NULL,
    service_type_code VARCHAR NOT NULL REFERENCES service_types(service_code),
    price_per_wheel INTEGER NOT NULL,
    price_per_set INTEGER NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT (NOW() AT TIME ZONE 'Europe/Moscow'),

    CONSTRAINT valid_rim_size CHECK (rim_size BETWEEN 12 AND 19),
    UNIQUE(rim_size, service_type_code)
);

CREATE TABLE available_slots (
    id BIGSERIAL PRIMARY KEY,
    date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT (NOW() AT TIME ZONE 'Europe/Moscow'),

    CONSTRAINT slots_time_check CHECK (end_time > start_time),
    UNIQUE(date, start_time, end_time)
);

CREATE TABLE bookings (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    user_phone VARCHAR,

    date VARCHAR NOT NULL,
    time VARCHAR NOT NULL,
    service VARCHAR NOT NULL,
    rim_radius INTEGER NOT NULL,
    total_price BIGINT,

    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS bookings CASCADE;
DROP TABLE IF EXISTS available_slots CASCADE;
DROP TABLE IF EXISTS prices CASCADE;
DROP TABLE IF EXISTS service_types CASCADE;