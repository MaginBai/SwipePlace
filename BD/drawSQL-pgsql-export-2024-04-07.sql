CREATE TABLE "Users"(
    "user_id" bigserial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "age" SMALLINT NOT NULL,
    "course" SMALLINT NOT NULL,
    "email" VARCHAR(320) UNIQUE NOT NULL,
    "login" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "date_registration" DATE NOT NULL
);

CREATE TABLE "Cards"(
    "card_id" bigserial PRIMARY KEY,
    "name" VARCHAR(320) NOT NULL,
    "author" BIGINT REFERENCES "Users"("user_id"),
    "address" VARCHAR(255) NOT NULL,
    "coordinates" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "img" VARCHAR(255) NOT NULL
);
