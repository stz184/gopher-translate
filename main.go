package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/stz184/gopher-translator/models"
	"gitlab.com/stz184/gopher-translator/server"
	"gitlab.com/stz184/gopher-translator/translator"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Gopher translator service!")
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	var wordRequest models.WordRequest
	err = json.Unmarshal(reqBody, &wordRequest)
	if err != nil {
		server.RespondWithError(w, "Please, provide a word")
		return
	}

	words := regexp.MustCompile(`\s+`).Split(wordRequest.Word, 2)
	if len(words) < 1 {
		server.RespondWithError(w, "Please, provide a non empty word.")
	}

	if len(words) > 1 {
		server.RespondWithError(w, "Please, provide a single word or use /sentence to translate sentences.")
		return
	}

	if wordRequest.Word == "" {
		server.RespondWithError(w, "Please, provide a word")
		return
	}

	translatedWord := translator.TranslateWord(wordRequest.Word)
	response := models.WordResponse{Word: translatedWord}
	server.ToJSON(w, response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePageHandler)
	myRouter.HandleFunc("/word", wordHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
