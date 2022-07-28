-- +migrate Up
create table if not exists posts
(
    id     text     primary key,
    title  text,
    url    text,
    poster text
);

create table if not exists votes
(
    id      text    primary key,
    post_id text,
    user    text
);

-- +migrate Down
drop table votes;
drop table posts;
