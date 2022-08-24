CREATE TABLE measure (
                         id SERIAL PRIMARY KEY,
                         name TEXT NOT NULL,
                         active BOOLEAN DEFAULT TRUE,
                         created_at timestamptz NOT NULL DEFAULT NOW(),
                         modified_at timestamptz NOT NULL DEFAULT NOW()
);

INSERT INTO measure (name)
VALUES
    ('grams'),
    ('ml'),
    ('litre'),
    ('tsp'),
    ('tbsp'),
    ('cup'),
    ('pinch'),
    ('as needed');