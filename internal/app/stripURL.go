package app

import (
	"log"
	"regexp"
)

var pattern *regexp.Regexp

func init() {
	var err error
	pattern, err = regexp.Compile("[^a-zA-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
}

func StripURL(url string, result chan string) {
	result <- pattern.ReplaceAllString(url, "")
}
