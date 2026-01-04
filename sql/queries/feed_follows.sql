-- name: CreateFeedFollow :one
with inserted_data as (
  insert into feed_follows(user_id, feed_id, created_at, updated_at)
  values (
    $1,
    $2,
    $3,
    $4
  )
  returning *
) select inserted_data.*, users.name as userame, feeds.name as feed_name
from inserted_data
inner join users on users.id = user_id
inner join feeds on feeds.id = feed_id;

-- name: GetFeedByUrl :one
select *
from feeds
where url = $1;

-- name: GetFeedFollowsForUser :many
select feeds.name as feed_name, users.name as user_name
from feed_follows
inner join feeds on feeds.id = feed_follows.feed_id
inner join users on users.id = feed_follows.user_id
where feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
delete
from feed_follows
where feed_follows.user_id = $1 and feed_follows.feed_id = $2;
