package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	go checkPrices(db)
	handler := NewHandler(db)
	srv := http.Server{
		Addr:    ":8000",
		Handler: routes(handler),
	}
	fmt.Println("Server Running on 8000")
	err = srv.ListenAndServe()
	log.Println(err)
}
