package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", connHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
	w.WriteHeader(http.StatusOK)
}
