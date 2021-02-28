package models

type WordResponse struct {
	Word string `json:"gopher-word"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
