package main

import (
	"math/rand"
	"time"
)

// json annotation
type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	Number     int64     `json:"number"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"createdAt"`
}

type TransferRequest struct {
	ToAccount int64 `json:"toAccount"`
	Amount    int   `json:"amount"`
}

// constructor, not going to use id random but the incremental id provided by Postgres
func NewAccount(firstName string, secondName string) *Account {
	return &Account{
		// ID:         rand.Intn(10000),
		FirstName:  firstName,
		SecondName: secondName,
		Number:     int64(rand.Intn(1000000)),
		CreatedAt:  time.Now().UTC(),
	}

}
