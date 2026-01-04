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

-- name: GetFeeds :many
select feeds.*, users.name as user_name
from feeds
inner join users
on feeds.user_id = users.id;