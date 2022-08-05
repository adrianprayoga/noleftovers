CREATE TABLE recipes (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT NOT NULL,
                       author INT
);