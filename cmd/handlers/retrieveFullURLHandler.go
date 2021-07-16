package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DeanFoleyDev/go-url-shortener/internal/app"
	"DeanFoleyDev/go-url-shortener/internal/data"
	"DeanFoleyDev/go-url-shortener/internal/db"
)

func RetrieveFullURLHandler(w http.ResponseWriter, r *http.Request) {
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

	err = app.ValidateShortURL(BaseURL, request.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	shortURL := app.StripURL(request.URL)

	longURL := db.Retrieve(shortURL)

	var response = data.Response{
		LongURL:  longURL,
		ShortURL: request.URL,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(jsonRes)
}
