CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at timestamptz NOT NULL DEFAULT NOW(),
                       modified_at timestamptz NOT NULL DEFAULT NOW()
);