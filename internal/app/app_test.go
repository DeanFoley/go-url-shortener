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
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			testURL := "https://df.dv/" + tt.testURL

			strippedURL := StripURL(testURL)

			if strippedURL != tt.expectedResult {
				t.Fatalf("URL has not been stripped properly: expected %s, got %s", tt.expectedResult, strippedURL)
			}
		})
	}
}

func Benchmark_StripURL(b *testing.B) {
	testURL := "https://df.dv/1234567890abcdef/"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StripURL(testURL)
	}
}

func Test_URLShortener(t *testing.T) {
	shortened := URLShortener()

	if len(shortened) > 16 || len(shortened) < 16 {
		t.Fatalf("Generated URL is wrong length: %v", len(shortened))
	}
}

func Benchmark_URLShortener(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		URLShortener()
	}
}

func Test_ValidateShortURL_Valid(t *testing.T) {
	testURL := "https://df.dv/1234567890abcdef/"

	baseURL := "https://df.dv/"

	err := ValidateShortURL(baseURL, testURL)

	if err != nil {
		t.Fatalf("Test failure: got %q", err)

	}
}

func Test_ValidateShortURL_Invalid(t *testing.T) {
	testURL := "https://goo.gl/1234567890abcdef/"

	baseURL := "https://df.dv/"

	err := ValidateShortURL(baseURL, testURL)

	if err == nil {
		t.Fatalf("Test failure: got %q", err)

	}
}

func Benchmark_ValidateURL(b *testing.B) {
	testURL := "https://df.dv/1234567890abcdef/"

	baseURL := "https://df.dv/"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ValidateShortURL(baseURL, testURL)
	}
}
