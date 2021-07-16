package handlers

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
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
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

	shortenedLink := app.URLShortener()

	go db.Store(request.URL, shortenedLink)

	var response = data.Response{
		LongURL:  request.URL,
		ShortURL: BaseURL + shortenedLink,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(jsonRes)
}
