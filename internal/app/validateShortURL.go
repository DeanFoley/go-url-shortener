package app

import (
	"fmt"
	"strings"
)

func ValidateShortURL(baseURL, url string) error {
	if !strings.Contains(url, baseURL) {
		return fmt.Errorf("URL is not valid for this service")
	}
	return nil
}
