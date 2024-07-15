package concurrency

type WebSiteChecker func(string) bool

type response struct {
	url    string
	result bool
}

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	resultCh := make(chan response, len(urls))

	for _, url := range urls {
		go func(u string) {
			resultCh <- response{u, wc(u)}
		}(url)
	}

	for range urls {
		r := <-resultCh
		result[r.url] = r.result
	}

	return result
}
