package app

func StripURL(url string) string {
	if url[len(url)-1] == '/' {
		return url[len(url)-17 : len(url)-1]
	}
	return url[len(url)-16:]
}
