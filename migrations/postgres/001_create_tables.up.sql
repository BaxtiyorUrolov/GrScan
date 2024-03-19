CREATE TYPE user_type_enum AS ENUM ('admin', 'customer', 'partner');

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    user_id SERIAL,
    phone VARCHAR(15) UNIQUE,
    login VARCHAR(15) UNIQUE,
    password VARCHAR(100),
    balance INT DEFAULT 0,
    count INT DEFAULT 0,
    key UUID UNIQUE,
    user_type user_type_enum NOT NULL,
    user_verify BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS verify_cods (
    id UUID PRIMARY KEY,
    phone VARCHAR(15) UNIQUE,
    code INT,
    created_at TIMESTAMP DEFAULT NOW()
);
