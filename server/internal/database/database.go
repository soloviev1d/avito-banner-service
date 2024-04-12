package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ParseData(dbUser, dbPass string) (string, error) {
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/avito_banners",
		dbUser,
		dbPass,
	)
	conn, err := pgx.Connect(context.Background(), url)
	for err != nil {
		conn, err = pgx.Connect(context.Background(), url)
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
