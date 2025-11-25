CREATE TABLE service_types (
    id BIGSERIAL PRIMARY KEY,
    service_code VARCHAR(20) NOT NULL UNIQUE,
    service_name VARCHAR(50) NOT NULL,
    description TEXT
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
    user_name VARCHAR,
    user_phone VARCHAR,
    slot_id BIGINT NOT NULL REFERENCES available_slots(id) ON DELETE RESTRICT,
    service_type_id BIGINT NOT NULL REFERENCES service_types(id) ON DELETE RESTRICT,

    rim_size INTEGER NOT NULL,
    wheel_count INTEGER NOT NULL DEFAULT 4,
    total_price INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'CONFIRMED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    CONSTRAINT unique_booking_slot UNIQUE(slot_id, chat_id),
    CONSTRAINT valid_wheel_count CHECK (wheel_count BETWEEN 1 AND 8)
);

CREATE INDEX idx_slots_date_available ON available_slots(date, is_available) WHERE is_available = true;
CREATE INDEX idx_slots_date_time ON available_slots(date, start_time);
CREATE INDEX idx_bookings_chat_id ON bookings(chat_id);
CREATE INDEX idx_bookings_slot_id ON bookings(slot_id);
CREATE INDEX idx_bookings_status ON bookings(status);
CREATE INDEX idx_prices_rim_service ON prices(rim_size, service_type_code) WHERE is_active = true;

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_bookings_updated_at BEFORE UPDATE ON bookings FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();