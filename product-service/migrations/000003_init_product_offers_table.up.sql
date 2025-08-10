CREATE TABLE product_offers (
    id SERIAL PRIMARY KEY,
    price DECIMAL NOT NULL,
    stock INT NOT NULL,
    is_active BOOLEAN NOT NULL,
    product_id BIGINT REFERENCES products(id),
    supplier_id BIGINT NOT NULL
);