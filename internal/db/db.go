package db

import "fmt"

var urlStore map[string]string

func init() {
	urlStore = make(map[string]string)
}

func Store(long string, short string, result chan bool) {
	if _, ok := urlStore[short]; ok {
		result <- false
	}
	urlStore[short] = long
	result <- true
}

type Response struct {
	LongURL string
	Err     error
}

func Retrieve(short string, result chan Response) {
	if longURL, ok := urlStore[short]; ok {
		result <- Response{
			LongURL: longURL,
			Err:     nil,
		}
		return
	}
	result <- Response{
		LongURL: "",
		Err:     fmt.Errorf("nothing found for key: %v", short),
	}
}
