CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    uid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    avatar VARCHAR(255),
    password VARCHAR(255),
    provider VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE chats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    sender_uid UUID NOT NULL,
    receiver_uid UUID NOT NULL,
    FOREIGN KEY (sender_uid) REFERENCES users(uid),
    FOREIGN KEY (receiver_uid) REFERENCES users(uid)
);