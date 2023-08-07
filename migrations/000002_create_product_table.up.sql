BEGIN;

CREATE TABLE IF NOT EXISTS products(
   id bigint unique not null primary key,
   name VARCHAR (200) UNIQUE NOT NULL,
   price NUMERIC(10, 2) NOT NULL,
   stock int DEFAULT '0' NULL
);

COMMIT;