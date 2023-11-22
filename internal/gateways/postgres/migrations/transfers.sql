CREATE TABLE IF NOT EXISTS transfers (
    id SERIAL PRIMARY KEY,
    account_origin_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    account_destination_id INT REFERENCES accounts(id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
