package db

var urlStore map[string]string

func init() {
	urlStore = make(map[string]string)
}

func Store(long string, short string) {
	urlStore[short] = long
}

func Retrieve(short string) string {
	return urlStore[short]
}
