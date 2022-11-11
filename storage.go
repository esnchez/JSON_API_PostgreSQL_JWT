package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	GetAccount() ([]*Account, error)

}

type PostgresStore struct {
	db *sql.DB
}

// Constructor and connection to db
func NewPostgresStorage() (*PostgresStore, error) {
	//default user and db created with docker run image
	connStr := "user=postgres dbname=postgres password=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

//Creating a table to the Postgres DB
func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {

	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		second_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err

}



//Storage interface implementation
func (s *PostgresStore) CreateAccount(acc *Account) error {
	
	query := `insert into account (
		first_name, second_name, number, balance, created_at)
		 values ( $1, $2, $3, $4, $5)`

	resp , err := s.db.Query(query, acc.FirstName, acc.SecondName, acc.Number, acc.Balance, acc.CreatedAt)
	if err != nil{
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) GetAccount() ([]*Account, error) {
	query := `select * from account`

	resp , err := s.db.Query(query)
	if err != nil{
		return nil, err
	}
	accounts := []*Account{}
	for resp.Next(){
		account := new(Account)
		if err := resp.Scan(
			&account.ID,
			&account.FirstName,
			&account.SecondName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		); err != nil {
			return nil,err
		}
		accounts = append(accounts, account )
	}
	return accounts, nil
}