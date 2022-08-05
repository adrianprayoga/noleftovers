CREATE TABLE ingredients (
                        id SERIAL PRIMARY KEY,
                        recipe_id INT NOT NULL,
                        position INT NOT NULL,
                        amount DECIMAL,
                        measure INT
);