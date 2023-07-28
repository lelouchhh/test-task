CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name VARCHAR(100) NOT NULL,
                       second_name VARCHAR(100),
                       last_name VARCHAR(100) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       debt NUMERIC CHECK (debt >= 0 AND debt <= 1000) DEFAULT 0
);