package entity

import (
	"github.com/google/uuid"
)

type TransactionStatus string

var (
	Pending  TransactionStatus = "pending"
	Approved TransactionStatus = "approved"
	Canceled TransactionStatus = "canceled"
	Refused  TransactionStatus = "refused"
)

type Transaction struct {
	Id        uuid.UUID
	Payer     uuid.UUID
	Payee     uuid.UUID
	Status    string
	Ammount   float64
	CreatedAt int64
	UpdatedAt int64
}
