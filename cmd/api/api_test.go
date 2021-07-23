package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"DeanFoleyDev/go-url-shortener/internal/data"
	"DeanFoleyDev/go-url-shortener/internal/db"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	stubMux := mux.NewRouter()

	stubMux.HandleFunc("/{short}", RetrieveFullURLHandler)
	stubMux.HandleFunc("shorten", ShortenURLHandler)

	stubServer := httptest.NewServer(stubMux)

	BaseURL = stubServer.URL

	os.Exit(m.Run())
}

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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Bad response from handler: %v", resp.StatusCode)
	}

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var result data.Response
	err = json.Unmarshal(output, &result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
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
		resp.Body.Close()
	}
}

func Test_RetrieveFullURLHandler(t *testing.T) {
	writer := httptest.NewRecorder()

	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"
	shortened := "1234567890abcdef"

	request := httptest.NewRequest("GET", fmt.Sprintf("%s/%s", BaseURL, shortened), nil)

	dbConfirmChan := make(chan bool, 1)
	db.Store(longURL, shortened, dbConfirmChan)
	<-dbConfirmChan
	close(dbConfirmChan)

	RetrieveFullURLHandler(writer, request)

	if writer.Code != http.StatusOK {
		t.Fatalf("Bad response from handler: %v", writer.Code)
	}

	response, err := ioutil.ReadAll(writer.Body)
	if err != nil {
		t.Fatal(err)
	}

	var result data.Response
	err = json.Unmarshal(response, &result)
	if err != nil {
		t.Fatalf("Error decoding body: %v", err)
	}

	if result.LongURL != longURL {
		t.Fatalf("Result mismatch, stored URL is: %v", result.LongURL)
	}
}

// Apparently this function just can't be benchmarked; I couldn't find a fix for this at all.
func Benchmark_RetrieveFullURLHandler(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(RetrieveFullURLHandler))
	defer ts.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"
	shortened := "1234567890abcdef"

	dbConfirmChan := make(chan bool, 1)
	db.Store(longURL, shortened, dbConfirmChan)
	<-dbConfirmChan
	close(dbConfirmChan)

	request := fmt.Sprintf("%s/%s", ts.URL, shortened)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		resp, err := cl.Get(request)
		if err != nil {
			b.Fatalf("Error making get request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b.Fatal()
		}
	}
}
