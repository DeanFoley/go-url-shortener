package app

import (
	"testing"
)

func Test_StripURL(t *testing.T) {
	tests := []struct {
		testName       string
		testURL        string
		expectedResult string
	}{
		{
			testName:       "No Trailing Slash",
			testURL:        "1234567890abcdef",
			expectedResult: "1234567890abcdef",
		},
		{
			testName:       "Trailing Slash",
			testURL:        "1234567890abcdef/",
			expectedResult: "1234567890abcdef",
		},
		{
			testName:       "Garbage",
			testURL:        "123456789/#0abcdef",
			expectedResult: "1234567890abcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			stripChan := make(chan string, 1)
			StripURL(tt.testURL, stripChan)
			strippedURL := <-stripChan
			close(stripChan)

			if strippedURL != tt.expectedResult {
				t.Fatalf("URL has not been stripped properly: expected %s, got %s", tt.expectedResult, strippedURL)
			}
		})
	}
}

func Benchmark_StripURL(b *testing.B) {
	testURL := "https://df.dv/1234567890abcdef/"

	stripChan := make(chan string, 1)

	go func() {
		for {
			<-stripChan
		}
	}()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StripURL(testURL, stripChan)
	}
}

func Test_URLShortener(t *testing.T) {
	shortenedChan := make(chan string, 1)
	URLShortener(shortenedChan)
	shortened := <-shortenedChan
	close(shortenedChan)

	if len(shortened) > 16 || len(shortened) < 16 {
		t.Fatalf("Generated URL is wrong length: %v", len(shortened))
	}
}

func Benchmark_URLShortener(b *testing.B) {
	shortenedChan := make(chan string, 1)
	go func() {
		for {
			<-shortenedChan
		}
	}()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		URLShortener(shortenedChan)
	}
}

func Test_ValidateShortURL_Valid(t *testing.T) {
	testURL := "https://df.dv/1234567890abcdef/"
	baseURL := "https://df.dv/"

	errorChan := make(chan error, 1)
	ValidateShortURL(baseURL, testURL, errorChan)

	if err := <-errorChan; err != nil {
		t.Fatalf("Test failure: got %q", err)
	}
	close(errorChan)
}

func Test_ValidateShortURL_Invalid(t *testing.T) {
	testURL := "https://goo.gl/1234567890abcdef/"
	baseURL := "https://df.dv/"

	errorChan := make(chan error, 1)
	ValidateShortURL(baseURL, testURL, errorChan)

	if err := <-errorChan; err == nil {
		t.Fatalf("Test failure: got %q", err)
	}
	close(errorChan)
}

func Benchmark_ValidateShortURL(b *testing.B) {
	testURL := "https://df.dv/1234567890abcdef/"
	baseURL := "https://df.dv/"

	errorChan := make(chan error, 1)
	go func() {
		for {
			<-errorChan
		}
	}()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ValidateShortURL(baseURL, testURL, errorChan)
	}
}
