-- name: CreateFeed :one
insert into feeds(created_at, updated_at, name, url, user_id)
values (
  $1,
  $2,
  $3,
  $4,
  $5
)
returning *;