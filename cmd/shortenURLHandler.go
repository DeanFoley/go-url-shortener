package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DeanFoleyDev/go-url-shortener/internal/app"
	"DeanFoleyDev/go-url-shortener/internal/data"
	"DeanFoleyDev/go-url-shortener/internal/db"
)

// POST /shorten/ { "url": string }
// Returns a shortened URL for a given long URL
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Incorrect REST method: %s, please use POST.", r.Method)
		return
	}

	var request data.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortenedURL := baseURL + app.URLShortener(request.URL)

	go db.Store(request.URL, shortenedURL)

	var response = data.Response{
		LongURL:  request.URL,
		ShortURL: shortenedURL,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(jsonRes)
}
