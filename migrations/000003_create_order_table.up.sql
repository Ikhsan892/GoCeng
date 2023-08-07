BEGIN;

CREATE TABLE IF NOT EXISTS shipping_3pls(
    id bigint unique not null primary key,
    name varchar(100) not null
);

CREATE TABLE IF NOT EXISTS users(
    id bigint unique not null primary key,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    address json NULL
);

CREATE TYPE enum_status AS ENUM ('UNPAID','PAID','PROCESSING','SHIPPING','DONE','RETURN','REFUND','CANCELLED','FRAUD');

CREATE TABLE IF NOT EXISTS orders(
   id bigint unique not null primary key,
   order_id varchar(50) not null,
   user_id bigint not null,
   payment_type varchar(100) not null,
   shipping_3pl varchar(100) not null,
   shipping_3pl_id bigint not null,
   total_price numeric(10,2) default '0' null,
   awb varchar(255) null,
   status enum_status default 'UNPAID' null,
   created_at timestamp default now() not null,
   updated_at timestamp default now() not null,
   deleted_at timestamp null,

   CONSTRAINT fk_shipping_3pl
       FOREIGN KEY(shipping_3pl_id)
           REFERENCES shipping_3pls(id),

   CONSTRAINT fk_users
       FOREIGN KEY(user_id)
           REFERENCES users(id)
);

COMMIT;