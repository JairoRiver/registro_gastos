-- name: CreateUserGroup :one
INSERT INTO user_group (
  user_id,
  group_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUserGroup :one
SELECT *
FROM user_group
WHERE id = $1 
LIMIT 1;

-- name: GetUserGroupByUser :one
SELECT *
FROM user_group
WHERE user_id = $1 
LIMIT 1;

-- name: GetUserGroupByGroup :one
SELECT *
FROM user_group
WHERE group_id = $1 
LIMIT 1;

-- name: ListUserGroups :many
SELECT *
FROM user_group
LIMIT $1
OFFSET $2;

-- name: DeleteUserGroup :exec
DELETE FROM user_group
WHERE id = $1;

-- name: UpdateUserGroup :one
UPDATE user_group
SET
  user_id = COALESCE(sqlc.narg(user_id), user_id),
  group_id = COALESCE(sqlc.narg(group_id), group_id),
  updated_at = NOW()
WHERE
  id = sqlc.arg(id)
RETURNING *;