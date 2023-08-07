BEGIN;

CREATE TABLE IF NOT EXISTS shipping_3pls(
     id bigint unique not null primary key,
     name varchar(100) not null
);

COMMIT;