package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"DeanFoleyDev/go-url-shortener/internal/app"
	"DeanFoleyDev/go-url-shortener/internal/data"
	"DeanFoleyDev/go-url-shortener/internal/db"
)

func TestMain(t *testing.T) {

}

// Accepts a URL to be shortened
func Test_ShortenURLHandler(t *testing.T) {
	request := data.Request{
		URL: "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel",
	}
	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/shorten/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()

	ShortenURLHandler(writer, req)

	resp := writer.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Bad response from handler: %v", resp.StatusCode)
	}
}

func Benchmark_ShortenURLHandler(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(ShortenURLHandler))
	defer ts.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	request := data.Request{
		URL: "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel",
	}
	body, err := json.Marshal(request)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		resp, err := cl.Post(ts.URL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			b.Fatalf("Error posting request: %v", err)
		}
		defer resp.Body.Close()
	}
}

func Test_RetrieveFullURLHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(RetrieveFullURLHandler))
	defer ts.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"
	shortened := app.URLShortener()

	db.Store(longURL, shortened)

	request := data.Request{
		URL: "https://df.dv/" + shortened + "/",
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cl.Post(ts.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Bad response from handler: %v", resp.StatusCode)
	}

	var result data.Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding body: %v", err)
	}
	resp.Body.Close()

	if result.LongURL != longURL {
		t.Fatalf("Result mismatch, stored URL is: %v", result.LongURL)
	}
}

func Benchmark_RetrieveFullURLHandler(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(RetrieveFullURLHandler))
	defer ts.Close()

	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"
	shortened := app.URLShortener()

	db.Store(longURL, shortened)

	request := data.Request{
		URL: "https://df.dv/" + shortened + "/",
	}

	body, err := json.Marshal(request)
	if err != nil {
		b.Fatal(err)
	}

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		resp, err := cl.Post(ts.URL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			b.Fatalf("Post: %v", err)
		}
		resp.Body.Close()
	}
}
