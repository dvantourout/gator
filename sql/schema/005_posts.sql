-- +goose Up
create table posts(
  id serial primary key,
  created_at timestamp not null,
  updated_at timestamp not null,
  title text not null,
  url text not null unique,
  description text not null,
  published_at timestamp not null,
  feed_id integer not null,
  foreign key (feed_id) references feeds(id)
);

-- +goose Down
drop table posts;