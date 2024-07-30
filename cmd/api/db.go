package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	dsn := fmt.Sprintf("host=postgres port=5432 user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
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
