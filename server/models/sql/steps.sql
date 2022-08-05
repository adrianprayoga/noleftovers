CREATE TABLE steps (
                        id SERIAL PRIMARY KEY,
                        recipe_id INT NOT NULL,
                        position INT NOT NULL,
                        text TEXT NOT NULL,
                        picture TEXT
);