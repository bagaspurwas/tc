package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "time"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/transactions/new", NewTransactionHandler).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
