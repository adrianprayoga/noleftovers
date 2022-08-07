CREATE TABLE steps (
                        id SERIAL PRIMARY KEY,
                        recipe_id INT NOT NULL,
                        position INT NOT NULL,
                        text TEXT NOT NULL,
                        picture TEXT,
                        created_at timestamptz NOT NULL DEFAULT NOW(),
                        modified_at timestamptz NOT NULL DEFAULT NOW()
);