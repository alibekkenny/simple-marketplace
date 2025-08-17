CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    quantity INT NOT NULL,
    price DECIMAL NOT NULL,
    product_offer_id BIGINT NOT NULL,
    order_id BIGINT REFERENCES orders(id)
);