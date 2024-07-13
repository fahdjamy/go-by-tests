package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countDownFrom = 3
	Boom          = "Boom"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
	duration time.Duration
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(d.duration)
}

func main() {
	CountDown(os.Stdout, DefaultSleeper{
		duration: countDownFrom * time.Second,
	})
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countDownFrom; i > 0; i-- {
		_, err := fmt.Fprintln(writer, i)
		sleeper.Sleep()
		if err != nil {
			fmt.Println(err)
		}
	}
	_, _ = fmt.Fprintln(writer, Boom)
}
