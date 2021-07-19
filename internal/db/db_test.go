package db

import (
	"strconv"
	"testing"
)

func Test_Store(t *testing.T) {
	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"
	shortURL := "1234567890abcdef"
	resultChan := make(chan bool, 1)
	Store(longURL, shortURL, resultChan)
	<-resultChan

	returnedFromStore, ok := urlStore[shortURL]

	if !ok {
		t.Fatal("URL was not properly stored.")
	}

	if returnedFromStore != longURL {
		t.Fatal("URL stored doesn't match URL requested.")
	}
}

func Benchmark_Store(b *testing.B) {
	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"

	storeChan := make(chan bool, 1)
	go func() {
		for {
			<-storeChan
		}
	}()

	for n := 0; n < b.N; n++ {
		Store(longURL, strconv.Itoa(n), storeChan)
	}
}

func Test_Retrieve(t *testing.T) {
	shortURL := "1234567890abcdef"
	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"

	urlStore[shortURL] = longURL

	resultChan := make(chan Response, 1)

	Retrieve(shortURL, resultChan)

	result := <-resultChan

	if result.LongURL != longURL {
		t.Fatalf("Result not stored properly: expected %s got %s", longURL, result)
	}
}

func Benchmark_Retrieve(b *testing.B) {
	shortURL := "1234567890abcdef"
	longURL := "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel"

	urlStore[shortURL] = longURL

	retrieveChan := make(chan Response, 1)
	go func() {
		for {
			<-retrieveChan
		}
	}()

	for n := 0; n < b.N; n++ {
		Retrieve(shortURL, retrieveChan)
	}
}
