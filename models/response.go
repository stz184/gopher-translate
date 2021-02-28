package models

type WordResponse struct {
	Word string `json:"gopher-word"`
}

type SentenceResponse struct {
	Sentence string `json:"gopher-sentence"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
