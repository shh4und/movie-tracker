BEGIN;

CREATE SCHEMA tracker;


CREATE TABLE IF NOT EXISTS tracker.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    profile_picture VARCHAR(255),
    status VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tracker.titles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    year VARCHAR(4),
    rated VARCHAR(10),
    released VARCHAR(20),
    runtime VARCHAR(10),
    genre VARCHAR(255),
    director VARCHAR(255),
    writer VARCHAR(255),
    actors VARCHAR(255),
    plot TEXT,
    language VARCHAR(50),
    country VARCHAR(50),
    awards VARCHAR(255),
    poster TEXT,
    imdb_rating VARCHAR(10),
    imdb_id VARCHAR(20),
    type VARCHAR(20),
    production VARCHAR(255),
    response VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tracker.comments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    title_id INTEGER NOT NULL REFERENCES tracker.titles(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tracker.user_favorites (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    title_id INTEGER NOT NULL REFERENCES tracker.titles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tracker.user_ratings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    title_id INTEGER NOT NULL REFERENCES tracker.titles(id) ON DELETE CASCADE,
    rating INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tracker.watch_later (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    title_id INTEGER NOT NULL REFERENCES tracker.titles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tracker.watched_movies (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    title_id INTEGER NOT NULL REFERENCES tracker.titles(id) ON DELETE CASCADE,
    watched_on TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

END;

