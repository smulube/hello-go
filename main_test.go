package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIndexReturnsOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response status not %v", "200", response.Code)
	}
}
