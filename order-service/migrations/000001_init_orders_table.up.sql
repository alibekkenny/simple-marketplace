CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    total_amount DECIMAL NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP,
    user_id BIGINT NOT NULL,
    payment_method VARCHAR NOT NULL,
    shipping_address VARCHAR NOT NULL
);