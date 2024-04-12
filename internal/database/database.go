package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	dbPass string
	dbUser string
	dbUrl  string
	conn   *pgx.Conn
)

func PrepareDatabase() {
	mustSetEnvs()
	createIfNotExists()
}

func mustSetEnvs() {
	dbUser = os.Getenv("POSTGRES_USER")
	dbPass = os.Getenv("POSTGRES_PASSWORD")
	if dbUser == "" || dbPass == "" {
		log.Fatalln("Environment variables weren't set properly")
	}
	dbUrl = fmt.Sprintf(
		"postgres://%s:%s@postgres_db:5432/postgres",
		dbUser,
		dbPass,
	)
}

func createIfNotExists() {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	secondsAccum := 0
	for err != nil {
		conn, err = pgx.Connect(context.Background(), dbUrl)
		if secondsAccum > 15 {
			log.Fatalf("Failed to connect to database: %s\n", err)
		}
		time.Sleep(time.Second)
		secondsAccum += 1
	}

	query, err := initQuery()
	if err != nil {
		log.Fatalf("Failed to read data to initialize database: %s\n", err)
	}
	if _, err = conn.Exec(context.Background(), query); err != nil {
		log.Fatalf("Failed to initialize database: %s\n", err)
	}
}

func initQuery() (string, error) {
	b, err := os.ReadFile("postgres/init.sql")
	if err != nil {
		return "", err
	}
	return string(b), err
}
