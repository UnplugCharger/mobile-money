-- name: CreatePayment :one
insert into payments(
customer_phone,
amount,
mpesa_reference,
business_id
)values(
$1,$2,$3,$4
)returning *;