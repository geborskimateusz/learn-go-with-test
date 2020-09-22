package selects

import (
	"fmt"
	"net/http"
	"time"
)

var timeout = 10 * time.Second

func Racer(firstURL, secURL string) (winner string, error error) {
	return ConfigurableRacer(firstURL, secURL, timeout)
}

/*
Racer which takes two URLs and "races" them
by hitting them with an HTTP GET and returning the URL
which returned first. If none of them return within 10 seconds
then it should return an error.
*/
func ConfigurableRacer(firstURL, secURL string, timeout time.Duration) (winner string, error error) {

	// The first one to send a value "wins"
	select {
	case <-ping(firstURL):
		return firstURL, nil
	case <-ping(secURL):
		return secURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Request for %s and %s timed out", firstURL, secURL)
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
