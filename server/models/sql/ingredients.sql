CREATE TABLE ingredients (
                        id SERIAL PRIMARY KEY,
                        recipe_id INT NOT NULL,
                        position INT NOT NULL,
                        name VARCHAR NOT NULL,
                        amount DECIMAL NOT NULL,
                        measure INT NOT NULL,
                        created_at timestamptz NOT NULL DEFAULT NOW(),
                        modified_at timestamptz NOT NULL DEFAULT NOW()
);