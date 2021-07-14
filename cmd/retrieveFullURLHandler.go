package main

import (
	"DeanFoleyDev/go-url-shortener/internal/data"
	"DeanFoleyDev/go-url-shortener/internal/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func retrieveFullURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "Incorrect REST method: %s, please use GET.", r.Method)
		return
	}

	var request data.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longURL := db.Retrieve(request.URL)

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
