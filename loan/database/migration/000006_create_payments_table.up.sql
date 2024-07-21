CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    product_type VARCHAR(255) NOT NULL,
    product_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMP NOT NULL,
    paid_at TIMESTAMP NULL
);

CREATE INDEX idx_payments_product ON payments(product_type);
CREATE INDEX idx_payments_product_id ON payments(product_id);
CREATE INDEX idx_payments_expired_at ON payments(expired_at);
CREATE INDEX idx_payments_paid_at ON payments(paid_at);
