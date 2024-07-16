package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns fastest url", func(t *testing.T) {
		fastServer := createDelayedServer(0)
		slowServer := createDelayedServer(15)

		defer fastServer.Close()
		defer slowServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		urls := []string{slowURL, fastURL}

		got, err := Racer(urls, DefaultTimeOut)

		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}

		if got != fastURL {
			t.Errorf("got %s, expected %s", got, fastURL)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverOne := createDelayedServer(10)
		serverTwo := createDelayedServer(15)
		serverThree := createDelayedServer(20)
		defer serverOne.Close()
		defer serverTwo.Close()
		defer serverThree.Close()

		urls := []string{serverOne.URL, serverTwo.URL, serverThree.URL}
		_, err := Racer(urls, 10*time.Millisecond)

		if err == nil {
			t.Errorf("got nil, expected error")
		}
	})
}

func TestWebsiteRacer(t *testing.T) {
	t.Run("get fastest url", func(t *testing.T) {
		fastServer := createDelayedServer(0)
		slowServer := createDelayedServer(15)
		defer fastServer.Close()
		defer slowServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		fastestURL := WebsiteRacer(slowURL, fastURL)
		if fastestURL != fastURL {
			t.Errorf("got %s, expected %s", fastURL, fastestURL)
		}
	})
}

func createDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
