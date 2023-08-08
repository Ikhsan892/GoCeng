
-- name: GetById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateCustomer :exec
INSERT INTO users(username,password,email,address) VALUES ($1,$2,$3,$4);

-- name: GetProductsByIds :many
SELECT id,name,price::float4,stock::bigint FROM products WHERE id = ANY($1::int[]);

-- name: CreateProduct :exec
INSERT INTO products(name,price,stock) VALUES ($1,$2::float8,$3::int); 