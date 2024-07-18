package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

var sleepFetchTime = 50 * time.Millisecond

// We need to test that we do not write any kind of response on the error case.
// Sadly httptest.ResponseRecorder doesn't have a way of figuring this out,
// so we'll have to roll our own spy to test for this. SpyResponseWriter
type SpyResponseWriter struct {
	written bool
}

func (w *SpyResponseWriter) Header() http.Header {
	w.written = true
	return http.Header{}
}

func (w *SpyResponseWriter) Write(b []byte) (int, error) {
	w.written = true
	return 0, errors.New("not implemented")
}

func (w *SpyResponseWriter) WriteHeader(statusCode int) {
	w.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	dataCh := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store cancelled")
				return
			default:
				// We are simulating a slow process where we build the result slowly by appending the string,
				// character by character in a goroutine
				time.Sleep(sleepFetchTime)
				result += string(c)
			}
		}
		dataCh <- result
	}()

	// use a select to wait for the goroutine to finish its work or for the cancellation to occur.
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-dataCh:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	data := "go lang"
	t.Run("server runs correctly", func(t *testing.T) {
		spyStore := &SpyStore{response: data, t: t}

		svr := Server(spyStore)
		req := httptest.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		if resp.Body.String() != data {
			t.Errorf("got %q, want %q", resp.Body.String(), data)
		}

		spyStore.assertWasNotCancelled()
	})

	t.Run("server expected to call cancel", func(t *testing.T) {
		spyStore := &SpyStore{response: data, t: t}
		svr := Server(spyStore)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(10*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		resp := &SpyResponseWriter{}

		svr.ServeHTTP(resp, req)

		if resp.written {
			t.Error("response was not cancelled")
		}
	})
}
