package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

//Constructor and connection to db
func NewPostgresStorage() (*PostgresStore, error){
	//default user and db created with docker run image
	connStr := "user=postgres dbname=postgres password=1234 sslmode=disable"
	db, err := sql.Open("postgres",connStr)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	},nil
}