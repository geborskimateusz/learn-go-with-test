package concurrency

// WebsiteChecker which takes a single URL and returns a boolean.
// This is used by the function to check all the websites.
type WebsiteChecker func(string) bool

// CheckWebsites that checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value
// - true for a good response, false for a bad response.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
