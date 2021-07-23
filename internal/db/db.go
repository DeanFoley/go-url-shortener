package db

var urlStore map[string]string

var cmds chan request

type request struct {
	instruction string
	longURL     string
	shortURL    string
	urlChan     chan string
	replyChan   chan bool
}

func init() {
	urlStore = make(map[string]string)

	cmds = make(chan request, 100)

	go func(urlStore map[string]string, cmds chan request) {
		for cmd := range cmds {
			switch cmd.instruction {
			case "shorten":
				if urlStore[cmd.shortURL] == "" {
					urlStore[cmd.shortURL] = cmd.longURL
					cmd.replyChan <- true
				} else {
					cmd.replyChan <- false
				}
			case "longen":
				if longURL := urlStore[cmd.shortURL]; longURL != "" {
					cmd.urlChan <- longURL
					cmd.replyChan <- true
				} else {
					cmd.urlChan <- ""
					cmd.replyChan <- false
				}
			}
		}
	}(urlStore, cmds)
}

func Store(long string, short string, result chan bool) {
	replyChan := make(chan bool, 1)
	request := request{
		instruction: "shorten",
		longURL:     long,
		shortURL:    short,
		replyChan:   replyChan,
	}
	cmds <- request
	result <- <-replyChan
	close(replyChan)
}

type Response struct {
	LongURL string
	Result  bool
}

func Retrieve(short string, result chan Response) {
	replyChan := make(chan bool, 1)
	urlChan := make(chan string, 1)
	request := request{
		instruction: "longen",
		shortURL:    short,
		replyChan:   replyChan,
		urlChan:     urlChan,
	}
	cmds <- request
	result <- Response{
		LongURL: <-urlChan,
		Result:  <-replyChan,
	}
	close(urlChan)
	close(replyChan)
}
