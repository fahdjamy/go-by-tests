package concurrency

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func webSiteCheckerSpy(url string) bool {
	if url == "https://google.com" {
		return false
	}
	return true
}

func slowWebSitesStub(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return webSiteCheckerSpy(url)
}

func TestCheckWebsites(t *testing.T) {
	t.Run("website checker", func(t *testing.T) {
		websites := []string{
			"https://google.com",
			"https://apple.com",
			"https://amzon.com",
		}

		response := CheckWebsites(webSiteCheckerSpy, websites)
		expected := map[string]bool{
			"https://google.com": false,
			"https://apple.com":  true,
			"https://amzon.com":  true,
		}

		if !reflect.DeepEqual(response, expected) {
			t.Errorf("got %v, expected %v", response, expected)
		}
	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "https://google.com" + strconv.Itoa(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebSitesStub, urls)
	}
}
