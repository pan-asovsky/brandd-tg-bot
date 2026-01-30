-- +goose Up
INSERT INTO service_types (service_code, service_name, is_composite) VALUES
    ('TAKE_IT_OUT', 'Съём-установка', true),
    ('TIRE_SERVICE', 'Шиномонтаж', true),
    ('BALANCING', 'Балансировка', true),
    ('COMPLEX', 'Комплекс', true),
    ('TAKE_AND_TIRE', 'Съём-установка и шиномонтаж', false),
    ('TAKE_AND_BALANCING', 'Съём-установка и балансировка', false),
    ('TIRE_AND_BALANCING', 'Шиномонтаж и балансировка', false);

INSERT INTO prices (rim_size, service_type_code, price_per_wheel, price_per_set)
SELECT
    rim_size,
    st.service_code,
    price_per_wheel,
    price_per_set
FROM
    service_types st,
    (VALUES
         (12, 200, 800),
         (13, 200, 800),
         (14, 225, 900),
         (15, 230, 920),
         (16, 255, 1020),
         (17, 260, 1040),
         (18, 260, 1040),
         (19, 300, 1200)
    ) AS prices(rim_size, price_per_wheel, price_per_set)
WHERE st.service_code = 'TAKE_IT_OUT';

INSERT INTO prices (rim_size, service_type_code, price_per_wheel, price_per_set)
SELECT
    rim_size,
    st.service_code,
    price_per_wheel,
    price_per_set
FROM
    service_types st,
    (VALUES
         (12, 230, 920),
         (13, 230, 920),
         (14, 235, 940),
         (15, 250, 1000),
         (16, 255, 1020),
         (17, 280, 1120),
         (18, 300, 1200),
         (19, 330, 1320)
    ) AS prices(rim_size, price_per_wheel, price_per_set)
WHERE st.service_code = 'TIRE_SERVICE';

INSERT INTO prices (rim_size, service_type_code, price_per_wheel, price_per_set)
SELECT
    rim_size,
    st.service_code,
    price_per_wheel,
    price_per_set
FROM
    service_types st,
    (VALUES
         (12, 250, 1000),
         (13, 250, 1000),
         (14, 250, 1020),
         (15, 250, 1000),
         (16, 300, 1200),
         (17, 300, 1200),
         (18, 350, 1400),
         (19, 350, 1400)
    ) AS prices(rim_size, price_per_wheel, price_per_set)
WHERE st.service_code = 'BALANCING';

INSERT INTO prices (rim_size, service_type_code, price_per_wheel, price_per_set)
SELECT
    rim_size,
    st.service_code,
    price_per_wheel,
    price_per_set
FROM
    service_types st,
    (VALUES
         (12, 680, 2450),
         (13, 680, 2450),
         (14, 705, 2550),
         (15, 730, 2600),
         (16, 810, 2900),
         (17, 830, 3000),
         (18, 900, 3250),
         (19, 980, 3550)
    ) AS prices(rim_size, price_per_wheel, price_per_set)
WHERE st.service_code = 'COMPLEX';

-- +goose Down
DELETE FROM prices;
DELETE FROM service_types;
ALTER TABLE service_types AUTO_INCREMENT = 1;
ALTER TABLE prices AUTO_INCREMENT = 1;