CREATE TABLE IF NOT EXISTS users (
  uuid VARCHAR(255) PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone
);