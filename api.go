package main

import (
	"./db"
	"net/http"
	"github.com/gorilla/mux"
)

func GetBlockchainHandler(http.ResponseWriter, *http.Request) {
	return
}

func RootHandler(http.ResponseWriter, *http.Request) {
	return
}

func NodeHandler(http.ResponseWriter, *http.Request) {
	return
}

func TransactionHandler(http.ResponseWriter, *http.Request) {
	return	
}

func serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/join", NodeHandler)
	r.HandleFunc("/chain", GetBlockchainHandler)
	r.HandleFunc("/transaction", TransactionHandler)
	http.Handle("/", r)
}
