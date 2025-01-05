package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "0512"
	dbname := "stock"

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err.Error())
	}

	err = conn.Ping()

	return conn, err
}
