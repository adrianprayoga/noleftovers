CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email TEXT UNIQUE NOT NULL,
                       full_name TEXT,
                       password_hash TEXT,
                       auth_method TEXT NOT NULL,
                       oauth_id TEXT NOT NULL,
                       picture TEXT NOT NULL,
                       created_at timestamptz NOT NULL DEFAULT NOW(),
                       modified_at timestamptz NOT NULL DEFAULT NOW(),
                       last_login timestamptz NOT NULL DEFAULT NOW()
);

