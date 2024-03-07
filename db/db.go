package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	connStr := "user=go_user dbname=go_learning password=123456789 host=192.168.100.100 port=32768 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
