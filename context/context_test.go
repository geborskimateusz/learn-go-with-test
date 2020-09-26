package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	// cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// s.cancelled = true
// }

func TestHandler(t *testing.T) {

	data := "Hello World"

	t.Run("happy path", func(t *testing.T) {
		store := &SpyStore{data, t}
		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf(`got "%s" want "%s`, res.Body.String(), data)
		}
		// if store.cancelled {
		// 	t.Errorf("store should not be cancelled")
		// }
	})

	t.Run("cancel work if request in cancelled", func(t *testing.T) {
		store := &SpyStore{data, t}
		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		// if !store.cancelled {
		// 	t.Errorf("store should be cancelled, not told to")
		// }
	})
}
