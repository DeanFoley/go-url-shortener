package data

// Incoming JSON request for a short URL
type Request struct {
	URL string `json:"url"`
}

// Outgoing JSON response containing a requested URL and its shortened version
type Response struct {
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}
