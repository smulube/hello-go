package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":5000", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")
	w.Write([]byte("hello!"))
}
