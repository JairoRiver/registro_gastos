-- name: CreateGroup :one
INSERT INTO "Groups" (
  user_id,
  name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetGroup :one
SELECT *
FROM "Groups" AS g
WHERE g.id = $1 
LIMIT 1;

-- name: GetGroupByUser :one
SELECT *
FROM "Groups" as g
WHERE g.user_id = $1 
LIMIT 1;

-- name: ListGroups :many
SELECT *
FROM "Groups" as g
LIMIT $1
OFFSET $2;

-- name: DeleteGroup :exec
DELETE FROM "Groups"
WHERE id = $1;

-- name: UpdateGroups :one
UPDATE "Groups"
SET
  name = COALESCE(sqlc.narg(name), name),
  updated_at = NOW()
WHERE
  id = sqlc.arg(id)
RETURNING *;