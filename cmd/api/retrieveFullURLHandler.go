package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"DeanFoleyDev/go-url-shortener/internal/app"
	"DeanFoleyDev/go-url-shortener/internal/data"
	"DeanFoleyDev/go-url-shortener/internal/db"
)

func RetrieveFullURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Incorrect REST method: %s, please use GET.", r.Method), http.StatusBadRequest)
		return
	}

	path, err := url.Parse(r.URL.String())
	if err != nil {
		return
	}

	request := path.Path
	if request == "" {
		http.Error(w, "No request found.", http.StatusBadRequest)
		return
	}

	shortenChan := make(chan string, 1)
	go app.StripURL(request, shortenChan)
	shortURL := <-shortenChan
	close(shortenChan)

	if len(shortURL) > 16 || len(shortURL) < 16 {
		http.Error(w, "No valid URL found.", http.StatusBadRequest)
	}

	longURLChan := make(chan db.Response, 1)
	go db.Retrieve(shortURL, longURLChan)
	longURL := <-longURLChan
	close(longURLChan)

	var response = data.Response{
		LongURL:  longURL.LongURL,
		ShortURL: shortURL,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(jsonRes)
}
