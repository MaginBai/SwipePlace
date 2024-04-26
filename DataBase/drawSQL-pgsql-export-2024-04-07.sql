CREATE TABLE users(
    id bigserial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age SMALLINT NOT NULL,
    courses SMALLINT NOT NULL,
    email VARCHAR(320) UNIQUE NOT NULL,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    date_registration DATE NOT NULL
);

CREATE TABLE cards(
    id bigserial PRIMARY KEY,
    name VARCHAR(320) NOT NULL,
    user_id BIGINT REFERENCES users(id),
    address VARCHAR(255) NOT NULL,
    coordinates VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    img VARCHAR(255) NOT NULL
);