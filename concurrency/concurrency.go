package main

type WebsiteChecker func(string) bool

type WebsiteCheckResult struct {
	url    string
	result bool
}

func CheckWebsites(checkerFn WebsiteChecker, urls []string) (results map[string]bool) {
	results = make(map[string]bool)
	resultChan := make(chan WebsiteCheckResult)
	for _, url := range urls {
		go func(u string) {
			resultChan <- WebsiteCheckResult{u, checkerFn(u)}
		}(url)
	}
	for range urls {
		r := <-resultChan
		results[r.url] = r.result
	}
	return
}
