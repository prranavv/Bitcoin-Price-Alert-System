package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DBConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDblifeTime = 5 * time.Minute

func ConnectDB() (*DB, error) {
	//Open a connection pool
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=pranav sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	//testing whether the databasea is connected
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	//setting constants for the database connections
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxLifetime(maxDblifeTime)
	//assigning the DB struct to the database
	DBConn.SQL = db
	return DBConn, nil
}
