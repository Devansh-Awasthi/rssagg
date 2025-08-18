-- +goose up
create table posts(
    id uuid primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    title text not null,
    description text,
    published_at timestamptz not null,
    url text not null unique,
    feed_id uuid not null references feeds(id) on delete cascade
);

-- +goose down
drop table posts;
