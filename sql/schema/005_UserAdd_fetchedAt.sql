-- +goose up
alter table feeds add column fetched_at timestamp;
-- +goose down
alter table feeds drop column fetched_at;