DROP TABLE IF EXISTS users, rooms, room_participants, room_messages, movies;
-- Таблица Пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(250) NOT NULL,
    role VARCHAR(30) NOT NULL,
    deleted boolean NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Таблица Комнат
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    is_public boolean NOT NULL,
    password VARCHAR(50) NOT NULL,
    deleted boolean NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Таблица Участников Комнаты
CREATE TABLE room_participants (
    room_id INT REFERENCES rooms(id),
    user_id INT REFERENCES users(id),
    role VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL,
    deleted boolean NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (room_id, user_id)
);
-- Таблица Чата в Комнате
CREATE TABLE room_messages (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES rooms(id),
    user_id INT REFERENCES users(id),
    message TEXT NOT NULL,
    deleted boolean NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Таблица Фильмов
CREATE TABLE movies (
    MovieID SERIAL PRIMARY KEY,
    MovieTitle VARCHAR(255) NOT NULL,
    MovieDescription TEXT,
    ReleaseYear INT
);
