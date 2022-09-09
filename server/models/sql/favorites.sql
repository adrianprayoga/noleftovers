CREATE TABLE favorites (
                        recipe_id INT NOT NULL,
                        user_id INT NOT NULL,
                        created_at timestamptz NOT NULL DEFAULT NOW(),
                        modified_at timestamptz NOT NULL DEFAULT NOW(),
                        PRIMARY KEY (recipe_id, user_id)
);