-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users 
WHERE user_id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = ?;

-- name: CreateUser :one
INSERT INTO users (
	first_name,
	last_name,
	email,
	password) 
VALUES 
	(?, ?, ?, ?) 
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
	first_name = ?,
	last_name = ?,
	email = ?
WHERE user_id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE user_id = ?;
