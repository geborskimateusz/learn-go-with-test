package concurrency

import "fmt"

// WebsiteChecker which takes a single URL and returns a boolean.
// This is used by the function to check all the websites.
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites that checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value
// - true for a good response, false for a bad response.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for i, url := range urls {
		go func(u string, i int) {
			fmt.Println("Send to channel with index:", i)
			resultChannel <- result{u, wc(u)}
		}(url, i)
	}

	fmt.Println("After concurrent processing")
	for i := 0; i < len(urls); i++ {
		fmt.Println("Receiving from channel..")
		result := <-resultChannel
		fmt.Printf("Received from channel %v\n", result)
		results[result.string] = result.bool
	}

	return results
}
