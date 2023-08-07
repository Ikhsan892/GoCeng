
-- name: GetById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateCustomer :exec
INSERT INTO users(username,password,email,address) VALUES ($1,$2,$3,$4);

-- name: GetProductsByIds :many
SELECT * FROM products WHERE id = ANY($1::int[]);

-- name: SaveOrder :exec
INSERT INTO orders(order_id,user_id,payment_type)