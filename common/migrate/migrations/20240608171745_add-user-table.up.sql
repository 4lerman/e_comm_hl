CREATE TYPE role_type AS ENUM ('admin', 'client');

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    fullName VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    address VARCHAR(50) NOT NULL,
    registerDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    userRole role_type NOT NULL
);