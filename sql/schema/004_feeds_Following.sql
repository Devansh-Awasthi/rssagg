-- +goose up
create table feeds_following(
    id uuid primary key,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    user_id uuid not null references users(id) on delete cascade,
    feed_id uuid not null references feeds(id) on delete cascade,
    unique(user_id, feed_id)
);
-- +goose down
drop table feeds_following;