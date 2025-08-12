-- +goose up
create table feeds(
    id uuid primary key,
    name text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    url text unique not null,
    user_id uuid references users(id) on delete cascade
);
-- +goose down
drop table feeds;