package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store fetches data from request
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server is simple implementation of Http flow
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return //log sth here
		}

		fmt.Fprint(w, data)
	}
}
