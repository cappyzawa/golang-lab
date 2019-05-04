package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiHnadler struct{}

func main() {
	mux := http.NewServeMux()
	registHandler(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello")
	return
}

func worldHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "world")
	return
}

func registHandler(mux *http.ServeMux) {
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/world", worldHandler)
}
