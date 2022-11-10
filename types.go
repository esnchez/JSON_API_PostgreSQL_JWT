package main

import "math/rand"

//json notation
type Account struct {
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	SecondName string `json:"secondName"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
}

func NewAccount(firstName string, secondName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		SecondName: secondName,
		Number: int64(rand.Intn(1000000)), 
	}

}



