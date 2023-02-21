-- name: CreateType :one
INSERT INTO "Type" (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetType :one
SELECT *
FROM "Type"
WHERE id = $1 
LIMIT 1;

-- name: ListTypes :many
SELECT *
FROM "Type";

-- name: DeleteType :exec
DELETE FROM "Type"
WHERE id = $1;

-- name: UpdateType :one
UPDATE "Type"
SET
  name = COALESCE(sqlc.narg(name), name),
  updated_at = NOW()
WHERE
  id = sqlc.arg(id)
RETURNING *;