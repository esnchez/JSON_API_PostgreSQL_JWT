package main

import (
	"log"
	"github.com/joho/godotenv"
)

func main() {
	//loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file..")
	}

	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NEWAPIServer(":3000", store)
	server.Run()
	

}
