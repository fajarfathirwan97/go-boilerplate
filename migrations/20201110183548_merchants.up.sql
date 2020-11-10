CREATE TABLE IF NOT EXISTS merchants (
    uuid VARCHAR(255) PRIMARY KEY,
    user_uuid VARCHAR(255) NOT NULL REFERENCES users (uuid) ON DELETE CASCADE,
    merchant_name VARCHAR(255) NOT NULL,
    merchant_address TEXT
);