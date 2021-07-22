package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	hostName string

	client = &http.Client{}
)

type response struct {
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}

type request struct {
	URL string `json:"URL"`
}

func main() {
	hostName = "http://localhost:8080"

	shorten := flag.NewFlagSet("shorten", flag.ExitOnError)
	shortenURL := shorten.String("shorten-URL", "", "shorten-URL")

	longen := flag.NewFlagSet("longen", flag.ExitOnError)
	longenURL := longen.String("longen-URL", "", "longen-URL")

	switch os.Args[1] {
	case "shorten":
		shorten.Parse(os.Args[2:])
		result, err := sendShortenRequest(*shortenURL)
		if err != nil {
			fmt.Printf("Error! Try again. Error was: %q\n", err)
			return
		}
		fmt.Println(result)
	case "longen":
		longen.Parse(os.Args[2:])
		result, err := sendLongenRequest(*longenURL)
		if err != nil {
			fmt.Printf("Error! Try again. Error was: %q\n", err)
			return
		}
		fmt.Println(result)
	default:
		fmt.Println("Error.")
	}
}

func sendShortenRequest(shortenURL string) (string, error) {
	body := request{
		URL: shortenURL,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	resp, err := client.Post(fmt.Sprintf("%s/shorten/", hostName), "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response response

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", err
	}

	return response.ShortURL, nil
}

func sendLongenRequest(longenURL string) (string, error) {
	resp, err := client.Get(fmt.Sprintf("%s/%s", hostName, longenURL))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response response

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", err
	}

	return response.LongURL, nil
}
