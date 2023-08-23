package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	// conexao := "user=root dbname=loja_go password=123 host-localhost sslmode=disble"
	conexao := "root:123@/loja_go"
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
