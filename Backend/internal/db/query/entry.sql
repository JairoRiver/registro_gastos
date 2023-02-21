-- name: CreateEntry :one
INSERT INTO "Entry" (
  user_id,
  group_id,
  type_id,
  name,
  use_day,
  amount,
  cost,
  cost_indicator,
  place
) VALUES (
  $1, $2, $3, $4, $5,$6, $7, $8, $9
) RETURNING *;

-- name: GetEntry :one
SELECT en.id
      ,en.user_id
      ,en.group_id
      ,en.use_day
      ,en.name 
      ,ty.name as type
      ,en.amount
      ,en.cost
      ,en.cost_indicator
      ,en.place
FROM "Entry" as en
JOIN "Type" as ty on ty.id = en.type_id
WHERE en.id = $1 
LIMIT 1;

-- name: ListEntryByUser :many
SELECT en.id
      ,en.user_id
      ,en.group_id
      ,en.use_day
      ,en.name 
      ,ty.name as type
      ,en.amount
      ,en.cost
      ,en.cost_indicator
      ,en.place
FROM "Entry" as en
JOIN "Type" as ty on ty.id = en.type_id
WHERE en.user_id = $1;

-- name: ListEntryByGroup :many
SELECT en.id
      ,en.user_id
      ,en.group_id
      ,en.use_day
      ,en.name 
      ,ty.name as type
      ,en.amount
      ,en.cost
      ,en.cost_indicator
      ,en.place
FROM "Entry" as en
JOIN "Type" as ty on ty.id = en.type_id
WHERE en.group_id = $1;

-- name: DeleteEntry :exec
DELETE FROM "Entry"
WHERE id = $1;

-- name: UpdateEntry :one
UPDATE "Entry"
SET
  user_id = COALESCE(sqlc.narg(user_id), user_id),
  updated_at = NOW(),
  group_id = COALESCE(sqlc.narg(group_id), group_id),
  type_id = COALESCE(sqlc.narg(type_id), type_id),
  name = COALESCE(sqlc.narg(name), name),
  use_day = COALESCE(sqlc.narg(use_day), use_day),
  amount = COALESCE(sqlc.narg(amount), amount),
  cost = COALESCE(sqlc.narg(cost), cost),
  cost_indicator = COALESCE(sqlc.narg(cost_indicator), cost_indicator),
  place = COALESCE(sqlc.narg(place), place)
WHERE
  id = sqlc.arg(id)
RETURNING *;