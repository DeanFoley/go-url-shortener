package main

import (
	"log"
	"net/http"
)

var (
	baseURL string
)

func main() {
	baseURL = "https://df.dv/"

	http.HandleFunc("/shorten/", shortenURLHandler)
	http.HandleFunc("/longen/", retrieveFullURLHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
