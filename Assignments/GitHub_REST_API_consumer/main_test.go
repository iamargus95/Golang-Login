package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetInput(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatalf("Test failed, Check link again.\nERROR: %v", err)
	}

	rec := httptest.NewRecorder()
	getInput(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.StatusCode)
	}
}
