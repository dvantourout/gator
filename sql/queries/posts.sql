-- name: CreatePost :one
insert into posts(
  created_at, updated_at, url, title, description, published_at, feed_id
) values (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) returning *;

-- name: GetPostsForUser :many
select posts.*
from posts
left join feed_follows on feed_follows.feed_id = posts.feed_id
where feed_follows.user_id = $1
order by posts.published_at desc
limit $2;