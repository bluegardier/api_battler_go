package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go_initial/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	// Faz um ping no banco para testar a conexão.
	err = conn.Ping()

	return conn, err
}