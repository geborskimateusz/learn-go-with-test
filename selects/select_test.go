package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("get faster url", func(t *testing.T) {
		slowServer := delayedRequest(20 * time.Millisecond)
		fastServer := delayedRequest(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns an error after 10s of no response", func(t *testing.T) {
		slowServer := delayedRequest(20 * time.Second)
		fastServer := delayedRequest(21 * time.Second)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowURL, fastURL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected error, got nill")
		}

	})
}

func delayedRequest(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
