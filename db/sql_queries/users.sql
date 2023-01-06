-- name: CreateUser :one
insert into users (
name, email, password
) values (
$1, $2, $3
) returning *;


-- name: GetUserByEmail :one
select * from users where email = $1 limit 1;

-- name: GetUserByID :one
select * from users where id = $1 limit 1;

-- name: GetUserByName :one
select * from users where name = $1 limit 1;

-- name: UpdateUserName :one
update users set name = $2 where id = $1 returning *;

-- name: UpdateUserEmail :one
update users set email = $2 where id = $1 returning *;

-- name: UpdateUserPassword :one
update users set password = $2 where id = $1 returning *;

-- name: DeleteUser :exec
delete from users where id = $1;

-- name: ListUsers :many
select * from users;