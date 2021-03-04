package main

import (
	"github.com/gorilla/mux"
	"github.com/stz184/gopher-translator/server"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", server.HomePageHandler)
	myRouter.HandleFunc("/word", server.WordHandler).Methods("POST")
	myRouter.HandleFunc("/sentence", server.SentenceHandler).Methods("POST")
	myRouter.HandleFunc("/history", server.HistoryHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
