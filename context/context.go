package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fetch, err := store.Fetch(ctx)

		if err != nil {
			return
		}

		_, err = fmt.Fprint(w, fetch)
		if err != nil {
			log.Println(err)
		}
	}
}
