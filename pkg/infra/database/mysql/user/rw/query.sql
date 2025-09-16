-- name: CreateOrder :exec
INSERT INTO orders (order_id, order_qty, amount)
VALUES (?, ?, ?);

-- name: GetOrderByID :one
SELECT order_id, order_qty, amount, created_at, updated_at
FROM orders
WHERE order_id = ?;

-- name: GetAllOrders :many
SELECT order_id, order_qty, amount, created_at, updated_at
FROM orders
ORDER BY created_at DESC;