package main

import "math/rand"

type Account struct {
	ID int
	FirstName string
	SecondName string
	Number int64
	Balance int64
}

func NewAccount(firstName string, secondName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		SecondName: secondName,
		Number: int64(rand.Intn(1000000)), 
	}

}



