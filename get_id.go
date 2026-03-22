package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	var id string
	email := "anugrahsputra@gmail.com"
	err = db.QueryRow("SELECT id FROM profiles WHERE email = $1", email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Profile not found")
			return
		}
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Printf("%s\n", id)
}
