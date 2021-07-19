package app

import (
	"fmt"
	"strings"
)

func ValidateShortURL(baseURL, url string, resultChan chan error) {
	if !strings.Contains(url, baseURL) {
		resultChan <- fmt.Errorf("URL is not valid for this service")
		return
	}
	resultChan <- nil
}
