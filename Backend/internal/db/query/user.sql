-- name: CreateUser :one
INSERT INTO "Users" (
  username,
  email,
  password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT *
FROM "Users"
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT u.id
      ,u.username
      ,u.password 
FROM "Users" AS u
WHERE u.username = $1
LIMIT 1;

-- name: ListUsers :many
SELECT u.id
      ,u.username
      ,u.email
      ,u.created_at
      ,u.updated_at
FROM "Users" AS u 
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1 CASCADE;

-- name: UpdateUser :one
UPDATE "Users"
SET
  password = COALESCE(sqlc.narg(password), password),
  updated_at = NOW(),
  username = COALESCE(sqlc.narg(username), username),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  id = sqlc.arg(id)
RETURNING id, username, updated_at;