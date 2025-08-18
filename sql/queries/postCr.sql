-- name: CreatePost :one
insert into posts (id,created_at,updated_at,title,description,published_at,url,feed_id) 
values($1,$2,$3,$4,$5,$6,$7,$8)
returning *;
-- name: GetPostforUser :many
select posts.* from posts join feeds_following on posts.feed_id=feeds_following.feed_id 
where feeds_following.user_id=$1 order by posts.published_at desc limit $2; 

