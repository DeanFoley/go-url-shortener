package main

import (
	"DeanFoleyDev/go-url-shortener/cmd/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten/", handlers.ShortenURLHandler)
	http.HandleFunc("/longen/", handlers.RetrieveFullURLHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
