package main

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
)

// Handler is a struct that has all handlers attaching to it
type Handler struct {
	Db        *DB
	IDCounter int
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// RequestData is the format of the Request
type RequestPostData struct {
	Price int `json:"price"`
}

type RequestDeleteData struct {
	AlertID int `json:"alertid"`
}

type Alert struct {
	AlertID int
	Price   int
	Status  string
}

type DB struct {
	SQL *sql.DB
}

// PriceResponse is a struct that respresents the BitCoin value in USD
type PriceResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}
