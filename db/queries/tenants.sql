-- name: CreateTenant :one
INSERT INTO tenants (name) 
VALUES 
	(?)
RETURNING *;
