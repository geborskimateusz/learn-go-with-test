package selects

import (
	"net/http"
	"time"
)

/*
Racer which takes two URLs and "races" them
by hitting them with an HTTP GET and returning the URL
which returned first. If none of them return within 10 seconds
then it should return an error.
*/
func Racer(firstURL, secURL string) (winner string) {
	startFirst := time.Now()
	http.Get(firstURL)
	firstDuration := time.Since(startFirst)

	startSec := time.Now()
	http.Get(secURL)
	secDuration := time.Since(startSec)

	if firstDuration > secDuration {
		return firstURL
	}
	return secURL
}
