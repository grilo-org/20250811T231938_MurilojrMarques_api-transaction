package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type postgresDB struct {
	Db *sql.DB
}

func NewPostgresDB() (*postgresDB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		Env.DbUser, Env.DbPassword, Env.DbServer, Env.DbPort, Env.DbDataBase)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado no banco!")
	return &postgresDB{Db: db}, nil
}
