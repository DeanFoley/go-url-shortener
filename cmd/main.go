package main

import (
	"log"
	"net/http"

	"DeanFoleyDev/go-url-shortener/cmd/api"
)

func main() {
	http.HandleFunc("/shorten/", api.ShortenURLHandler)
	http.HandleFunc("/", api.RetrieveFullURLHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
