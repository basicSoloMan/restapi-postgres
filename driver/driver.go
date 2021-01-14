package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host	 = "localhost"
	port	 = 5432
	user 	 = "postgres"
	password = "root"
	dbname	 = "Books"
)

var db *sql.DB
var err error

func atTheDisco(err error) {
	if err != nil {
		panic(err)
	}
}

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	atTheDisco(err)

	err = db.Ping()
	atTheDisco(err)

	return db
}