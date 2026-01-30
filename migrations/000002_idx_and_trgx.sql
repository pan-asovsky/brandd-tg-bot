-- +goose Up
CREATE UNIQUE INDEX unique_active_booking_slot
    ON bookings (date, time, chat_id)
    WHERE is_active;

CREATE INDEX idx_slots_date_available ON available_slots(date, is_available) WHERE is_available = true;
CREATE INDEX idx_slots_date_time ON available_slots(date, start_time);
CREATE INDEX idx_bookings_chat_id ON bookings(chat_id);
CREATE INDEX idx_bookings_status ON bookings(status);
CREATE INDEX idx_prices_rim_service ON prices(rim_size, service_type_code) WHERE is_active = true;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_bookings_updated_at
    BEFORE UPDATE ON bookings
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS update_bookings_updated_at ON bookings;
DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS idx_slots_date_available;
DROP INDEX IF EXISTS idx_slots_date_time;
DROP INDEX IF EXISTS idx_bookings_chat_id;
DROP INDEX IF EXISTS idx_bookings_status;
DROP INDEX IF EXISTS idx_prices_rim_service;
DROP INDEX IF EXISTS unique_active_booking_slot;