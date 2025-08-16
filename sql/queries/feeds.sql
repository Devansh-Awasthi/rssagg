-- name: CreateFeed :one
insert into feeds (id,name,created_at,updated_at,url,user_id) 
values($1,$2,$3,$4,$5, $6
)
returning *;
-- name: GetFeeds :many
select * from feeds ;
-- name: GetNextFeedsToFetch :many
select * from feeds order by fetched_at asc nulls first limit $1; 
-- name: UpdateNextFeedToFetch :one
update feeds set fetched_at=now() ,updated_at=now() where id=$1
returning *;
