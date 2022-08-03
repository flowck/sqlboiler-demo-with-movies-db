-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    is_adult BOOLEAN NOT NULL,
    budget FLOAT,
    homepage VARCHAR(150),
    imdb_id VARCHAR(15),
    original_title VARCHAR(150),
    original_language VARCHAR(50),
    overview TEXT,
    popularity FLOAT,
    poster_path VARCHAR(150),
    release_date DATE,
    revenue FLOAT,
    runtime FLOAT,
    tagline VARCHAR(150),
    title VARCHAR(150),
    video_url VARCHAR(150),
    vote_average FLOAT,
    vote_count INT
);

CREATE TABLE IF NOT EXISTS genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE
);

CREATE TABLE IF NOT EXISTS movie_genres (
    id SERIAL PRIMARY KEY,
    movie_id SERIAL,
    genre_id SERIAL,

    CONSTRAINT fk_movies
        FOREIGN KEY (movie_id)
        REFERENCES movies (id)
        ON DELETE CASCADE,

    CONSTRAINT fk_genres
        FOREIGN KEY (genre_id)
        REFERENCES genres (id)
        ON DELETE SET NULL
);

-- CREATE TABLE IF NOT EXISTS
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movie_genres;
DROP TABLE genres;
DROP TABLE movies;
-- +goose StatementEnd
