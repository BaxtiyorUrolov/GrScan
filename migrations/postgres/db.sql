CREATE TABLE users (
    id UUID PRIMARY KEY,
    user_id SERIAL,
    phone VARCHAR(15),
    login VARCHAR(30),
    password VARCHAR(70),
    balance MONEY,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gr_users (
    id UUID PRIMARY KEY,
    user_id INT PRIMARY KEY,
    username VARCHAR(40),
    nickname VARCHAR(50)
);

CREATE TABLE groups (
    id UUID PRIMARY KEY,
    group_id INT PRIMARY KEY,
    username VARCHAR(40),
    name VARCHAR(50),
    user_id UUID REFERENCES gr_users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
