package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestToJSON(t *testing.T) {
	w := httptest.NewRecorder()
	input := struct {
		Foo string `json:"foo"`
	}{Foo: "bar"}
	output := "{\"foo\":\"bar\"}\n"

	ToJSON(w, input)

	if w.Code != http.StatusOK {
		t.Log("Expected status code 200 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != output {
		t.Log("Expected JSON output", output, " but got", w.Body.String())
		t.Fail()
	}
}

func TestRespondWithError(t *testing.T) {
	w := httptest.NewRecorder()
	input := "Very helpful error message"
	output := "{\"error\":\"" + input + "\"}\n"

	RespondWithError(w, input)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Log("Expected status code 400 but got", w.Code)
		t.Fail()
	}

	if w.Body.String() != output {
		t.Log("Expected JSON output", output, "but got", w.Body.String())
		t.Fail()
	}
}
