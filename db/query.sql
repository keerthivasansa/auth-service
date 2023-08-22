-- name: CreateUser :exec
INSERT INTO users (id, provider, providerId, name, password) VALUES (?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE provider = ? AND providerId = ? LIMIT 1;