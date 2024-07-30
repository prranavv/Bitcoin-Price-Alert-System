package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DBConn = &DB{}

func openDB(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	DBConn.SQL = db
	return DBConn, nil
}

func ConnectDB() (*DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=pranav dbname=postgres sslmode=disable timezone=UTC connect_timeout=5"
	counts := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready....")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection, nil
		}
		if counts > 10 {
			log.Println(err)
			return nil, err
		}
		log.Println("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
