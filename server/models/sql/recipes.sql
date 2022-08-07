CREATE TABLE recipes (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT NOT NULL,
                       author INT,
                       created_at timestamptz NOT NULL DEFAULT NOW(),
                       modified_at timestamptz NOT NULL DEFAULT NOW()
);