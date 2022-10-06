package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	conexao := "user=postgres dbname=tfalc_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
