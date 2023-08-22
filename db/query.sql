-- name: CreateUser :exec
INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?);

-- name: GetUserWithEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;