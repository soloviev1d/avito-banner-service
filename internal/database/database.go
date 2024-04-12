package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func ParseData(dbUser, dbPass string) (string, error) {
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/avito_banners",
		dbUser,
		dbPass,
	)
	conn, err := pgx.Connect(context.Background(), url)
	// wait for postgres to start
	for err != nil {
		fmt.Println(err)
		conn, err = pgx.Connect(context.Background(), url)
		time.Sleep(1)
	}

	var (
		id     int64
		data   string
		result string
	)
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM test_table;",
	)
	for rows.Next() {
		rows.Scan(&id, &data)
		result += fmt.Sprintf("%s(%d)", data, id)
	}

	return result, nil
}
