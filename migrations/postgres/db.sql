CREATE TYPE user_type_enum AS ENUM ('admin', 'customer', 'partner');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    user_id SERIAL,
    phone VARCHAR(15) UNIQUE,
    login VARCHAR(15) UNIQUE,
    password VARCHAR(100),
    balance INT DEFAULT 0,
    count INT DEFAULT 0,
    key UUID UNIQUE,
    user_type user_type_enum NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
