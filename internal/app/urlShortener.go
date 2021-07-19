package app

import (
	"math/rand"
	"time"
)

var seed *rand.Rand

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"1234567890"

func init() {
	seed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func URLShortener(result chan string) {
	url := make([]byte, 16)

	for i := range url {
		url[i] = charset[seed.Intn(len(charset))]
	}

	result <- string(url)
}
