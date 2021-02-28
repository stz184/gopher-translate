package models

type WordRequest struct {
	Word string `json:"english-word"`
}

type SentenceRequest struct {
	Sentence string `json:"english-sentence"`
}
