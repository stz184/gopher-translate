package server

import (
	"encoding/json"
	"gitlab.com/stz184/gopher-translator/models"
	"gitlab.com/stz184/gopher-translator/translator"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
)

var history = make(map[string]string)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	var err = "Invalid endpoint. Please use POST /word to translate a word, POST /sentence to translate a sentence, or " +
		"GET /history to view the translations history"
	RespondWithError(w, err)
}

func WordHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, "Error reading the request body")
		return
	}

	var wordRequest models.WordRequest
	err = json.Unmarshal(reqBody, &wordRequest)
	if err != nil {
		RespondWithError(w, "Invalid payload")
		return
	}

	if wordRequest.Word == "" {
		RespondWithError(w, "Please, provide a word")
		return
	}

	words := regexp.MustCompile(`\s+`).Split(wordRequest.Word, 2)
	if len(words) > 1 {
		RespondWithError(w, "Please, provide a single word or use /sentence to translate sentences.")
		return
	}

	translatedWord := translator.TranslateWord(wordRequest.Word)
	history[wordRequest.Word] = translatedWord

	response := models.WordResponse{Word: translatedWord}
	ToJSON(w, response)
}

func SentenceHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, "Error reading the request body")
		return
	}

	var sentenceRequest models.SentenceRequest
	err = json.Unmarshal(reqBody, &sentenceRequest)
	if err != nil {
		RespondWithError(w, "Invalid payload")
		return
	}
	if sentenceRequest.Sentence == "" {
		RespondWithError(w, "Please, provide a sentence.")
		return
	}

	words := regexp.MustCompile(`\s+`).Split(sentenceRequest.Sentence, 2)
	if len(words) == 1 {
		RespondWithError(w, "If you need to translate a single word use /word instead.")
		return
	}

	translatedSentence := translator.TranslateSentence(sentenceRequest.Sentence)
	history[sentenceRequest.Sentence] = translatedSentence

	response := models.SentenceResponse{Sentence: translatedSentence}
	ToJSON(w, response)
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	var historyLen = len(history)
	var historyJSON = make([]map[string]string, 0, historyLen)

	englishInput := make([]string, 0, historyLen)
	for k := range history {
		englishInput = append(englishInput, k)
	}

	sort.Strings(englishInput)

	for _, k := range englishInput {
		translation := make(map[string]string, 1)
		translation[k] = history[k]
		historyJSON = append(historyJSON, translation)
	}

	ToJSON(w, historyJSON)
}
