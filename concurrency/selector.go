package concurrency

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const DefaultTimeOut = 10 * time.Second

func Racer(urls []string, timeout time.Duration) (string, error) {

	winner := ""
	leastTime := -1 * time.Millisecond

	for _, url := range urls {
		select {
		case elapsed := <-websitePing(url):
			if leastTime > elapsed || leastTime == -1*time.Millisecond {
				leastTime = elapsed
				winner = url
			}
		case <-time.After(timeout):
			return "", errors.New("")
		}
	}

	return winner, nil
}

func websitePing(url string) chan time.Duration {
	ch := make(chan time.Duration)
	go func() {
		start := time.Now()
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching ", url, ": ", err)
		}
		elapsed := time.Since(start)
		ch <- elapsed
	}()
	return ch
}

func WebsiteRacer(urlOne, urlTwo string) string {
	// select allows you to wait on multiple channels.
	// The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(urlOne):
		return urlOne
	case <-ping(urlTwo):
		return urlTwo
	}
}

/*
Why struct{} and not another type like a bool? Well,
a chan struct{} is the smallest data type available from a memory perspective,
so we get no allocation versus a bool.
Since we are closing and not sending anything on the chan, why allocate anything?
*/
func ping(url string) chan struct{} {
	// Always make channels
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching ", url, ": ", err)
		}
		close(ch)
	}()
	return ch
}
