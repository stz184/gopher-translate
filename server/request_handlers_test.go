package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	HomePageHandler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Log("Expected status code 400 but got", w.Code)
		t.Fail()
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Log("Expected content type application/json but got", w.Header().Get("Content-Type"))
		t.Fail()
	}

	var output = "Invalid endpoint. Please use POST /word to translate a word, POST /sentence to translate a sentence, or " +
		"GET /history to view the translations history"

	if !strings.Contains(w.Body.String(), output) {
		t.Log("Expected output to contains", output, " but got", w.Body.String())
		t.Fail()
	}
}

func TestWordHandler(t *testing.T) {
	requests := []struct {
		Payload            *strings.Reader
		ExpectedStatusCode int
		ExpectedOutput     string
	}{
		{
			strings.NewReader(""),
			http.StatusBadRequest,
			"Invalid payload",
		},
		{
			strings.NewReader("{\"english-word\": \"\"}"),
			http.StatusBadRequest,
			"Please, provide a word",
		},
		{
			strings.NewReader("{\"english-word\": \"foo bar\"}"),
			http.StatusBadRequest,
			"Please, provide a single word or use /sentence to translate sentences.",
		},
		{
			strings.NewReader("{\"english-word\": \"apple\"}"),
			http.StatusOK,
			"gapple",
		},
	}

	for _, request := range requests {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/word", request.Payload)

		WordHandler(w, r)

		if w.Code != request.ExpectedStatusCode {
			t.Log("Expected status code 400 but got", w.Code)
			t.Fail()
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Log("Expected content type application/json but got", w.Header().Get("Content-Type"))
			t.Fail()
		}

		if !strings.Contains(w.Body.String(), request.ExpectedOutput) {
			t.Log("Expected output to contains", request.ExpectedOutput, " but got", w.Body.String())
			t.Fail()
		}
	}
}

func TestSentenceHandler(t *testing.T) {
	requests := []struct {
		Payload            *strings.Reader
		ExpectedStatusCode int
		ExpectedOutput     string
	}{
		{
			strings.NewReader(""),
			http.StatusBadRequest,
			"Invalid payload",
		},
		{
			strings.NewReader("{\"english-sentence\": \"\"}"),
			http.StatusBadRequest,
			"Please, provide a sentence",
		},
		{
			strings.NewReader("{\"english-sentence\": \"foo\"}"),
			http.StatusBadRequest,
			"If you need to translate a single word use /word instead.",
		},
		{
			strings.NewReader("{\"english-sentence\": \"I want an apple for breakfast\"}"),
			http.StatusOK,
			"GI antwogo gan gapple orfogo eakfastbrogo",
		},
	}

	for _, request := range requests {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/sentence", request.Payload)

		SentenceHandler(w, r)

		if w.Code != request.ExpectedStatusCode {
			t.Log("Expected status code 400 but got", w.Code)
			t.Fail()
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Log("Expected content type application/json but got", w.Header().Get("Content-Type"))
			t.Fail()
		}

		if !strings.Contains(w.Body.String(), request.ExpectedOutput) {
			t.Log("Expected output to contains", request.ExpectedOutput, " but got", w.Body.String())
			t.Fail()
		}
	}
}

func TestHistoryHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/history", nil)

	mockHistory := make(map[string]string)
	mockHistory["foo"] = "oofogo"
	mockHistory["apple"] = "gapple"
	mockHistory["I draw a square"] = "I draw a square"
	history = mockHistory
	HistoryHandler(w, r)

	if w.Code != http.StatusOK {
		t.Log("Expected status code 200 but got", w.Code)
		t.Fail()
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Log("Expected content type application/json but got", w.Header().Get("Content-Type"))
		t.Fail()
	}

	output := "[{\"I draw a square\":\"I draw a square\"},{\"apple\":\"gapple\"},{\"foo\":\"oofogo\"}]\n"
	if w.Body.String() != output {
		t.Log("Expected output to contains", output, " but got", w.Body.String())
		t.Fail()
	}
}
