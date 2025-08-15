-- name: CreateFeedFollow :one
insert into feeds_following(id,created_at,updated_at,user_id,feed_id) 
values($1,$2,$3,$4,$5)
returning *;
-- name: GetFeedFollow :many
select * from feeds_following where user_id=$1;
-- name: DeleteFeedFollow :exec
delete from feeds_following where id=$1 and user_id=$2;