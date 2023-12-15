package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitConnection() *sql.DB {
	connection := "user=catalog_user dbname=catalog_db password=catalog_pass host=localhost port=5435 sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateTables() {
	conn := InitConnection()
	_, err := conn.Query("create table IF NOT EXISTS products (id serial primary key, name varchar, description varchar,	price decimal, quantity integer)")
	if err != nil {
		panic(err.Error())
	}
}
