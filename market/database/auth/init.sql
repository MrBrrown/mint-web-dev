CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role_id INTEGER NOT NULL REFERENCES roles(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roles (name) VALUES
    ('admin'),
    ('manager'),
    ('user');

INSERT INTO users (username, email, password_hash, role_id)
VALUES (
    'admin',
    'admin@example.com',
    '$2a$10$O1Cb.oan.9tkbuG8y9cqnOoVSchooDohWv9Y8B.MKwZ8ti9YDy8au', -- пароль: admin
    1
);
