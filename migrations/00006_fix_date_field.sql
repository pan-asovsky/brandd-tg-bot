-- +goose Up
ALTER TABLE bookings
    ALTER COLUMN date TYPE date USING to_date(date, 'YYYY-MM-DD');

-- +goose Down
ALTER TABLE bookings
    ALTER COLUMN date TYPE varchar;