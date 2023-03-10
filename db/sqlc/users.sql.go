// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
insert into users (
name, email, hashed_password
) values (
$1, $2, $3
) returning id, name, email, hashed_password, date_created, date_modified, account_balance
`

type CreateUserParams struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateModified,
		&i.AccountBalance,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, name, email, hashed_password, date_created, date_modified, account_balance from users where email = $1 limit 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateModified,
		&i.AccountBalance,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select id, name, email, hashed_password, date_created, date_modified, account_balance from users where id = $1 limit 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateModified,
		&i.AccountBalance,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
select id, name, email, hashed_password, date_created, date_modified, account_balance from users where name = $1 limit 1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateModified,
		&i.AccountBalance,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
select id, name, email, hashed_password, date_created, date_modified, account_balance from users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.HashedPassword,
			&i.DateCreated,
			&i.DateModified,
			&i.AccountBalance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserEmail = `-- name: UpdateUserEmail :one
update users set email = $2 where id = $1 returning id, name, email, hashed_password, date_created, date_modified, account_balance
`

type UpdateUserEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmail, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateModified,
		&i.AccountBalance,
	)
	return i, err
}

const updateUserName = `-- name: UpdateUserName :one
update users set name = $2 where id = $1 returning id, name, email, hashed_password, date_created, date_modified, account_balance
`

type UpdateUserNameParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserName, arg.ID, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.DateCreated,
		&i.DateModified,
		&i.AccountBalance,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
update users set hashed_password = $2 where id = $1 returning id
`

type UpdateUserPasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	var id int64
	err := row.Scan(&id)
	return id, err
}
