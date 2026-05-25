package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComOBancoDeDados() *sql.DB {
	cenexao := "user=postgres dbname=alura_loja password=**** host=localhost sslmode=disable"
	db, err := sql.Open("postgres", cenexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
