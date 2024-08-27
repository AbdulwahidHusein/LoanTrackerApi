package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanStatus string

const (
	Pending  LoanStatus = "pending"
	Rejected LoanStatus = "rejected"
	Approved LoanStatus = "approved"
)

type Loan struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IssuerId primitive.ObjectID `json:"issuer_id" bson:"issuer_id"`
	Status   LoanStatus         `json:"status" bson:"status"`
	Amount   int                `json:"amount" bson:"amount"`
}

type LoanFilter struct {
	Status  string   `json:"status" bson:"status"`
	OrderBy []string `json:"order_by" bson:"order_by"`
}
