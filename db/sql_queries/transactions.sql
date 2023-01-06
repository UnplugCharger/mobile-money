-- name: CreateTransaction :one
insert into transactions (
from_user,
to_user,
amount
) values (
$1,$2,$3
) returning *;


-- name: GetTransaction :one
select * from transactions where id = $1;

-- name: ListTransactions :many
select * from transactions where from_user = $1 or to_user = $2 order by created_at desc limit $3 offset $4;

