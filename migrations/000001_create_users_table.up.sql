BEGIN;

CREATE TABLE IF NOT EXISTS users(
    id bigint unique not null primary key,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    address json NULL
);


COMMIT;