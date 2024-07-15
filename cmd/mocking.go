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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func main() {
	configurableSleeper := ConfigurableSleeper{
		duration: countDownFrom * time.Second,
		sleep:    time.Sleep,
	}
	CountDown(os.Stdout, configurableSleeper)
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

func (cs ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}
