-- +goose up
create table users(
    id uuid primary key,
    name text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);
-- +goose down
drop table users;