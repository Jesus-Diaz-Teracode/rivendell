package main

import (
	"log"

	"github.com/jedi4z/rivendell/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.CreateDatabase()
	if err != nil {
		log.Fatalf("Error connection with database %v", err)
	}

	r := config.CreateEngine(db)
	r.Run()
}
