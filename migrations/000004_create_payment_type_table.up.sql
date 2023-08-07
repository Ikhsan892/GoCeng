BEGIN;

CREATE TABLE IF NOT EXISTS payment_types(
     id bigint unique not null primary key,
     name varchar(100) not null
);

COMMIT;