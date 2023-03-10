// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID             int64          `json:"id"`
	CustomerPhone  string         `json:"customer_phone"`
	Amount         sql.NullInt64  `json:"amount"`
	Status         sql.NullString `json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	MpesaReference string         `json:"mpesa_reference"`
	FailureReason  sql.NullString `json:"failure_reason"`
	BusinessID     int64          `json:"business_id"`
}

type Transaction struct {
	ID          int64          `json:"id"`
	PaymentID   sql.NullInt64  `json:"payment_id"`
	PaymentType sql.NullString `json:"payment_type"`
	Amount      sql.NullInt64  `json:"amount"`
	FromUser    int64          `json:"from_user"`
	ToUser      int64          `json:"to_user"`
}

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	DateCreated    time.Time `json:"date_created"`
	DateModified   time.Time `json:"date_modified"`
	AccountBalance int64     `json:"account_balance"`
}
