package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Called     int
	Operations []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Operations = append(s.Operations, sleep)
	s.Called++
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Operations = append(s.Operations, write)
	return
}

func TestCountDown(t *testing.T) {

	t.Run("printing as expected", func(t *testing.T) {
		buff := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		CountDown(buff, spySleeper)
		expected := `3
2
1
Boom
`
		if buff.String() != expected {
			t.Errorf("CountDown(...) returned %q, want %q", buff.String(), expected)
		}
		if spySleeper.Called != 3 {
			t.Errorf("expected called to be 3, was called %d", spySleeper.Called)
		}
	})

	t.Run("operations called in expected order", func(t *testing.T) {
		spySleepWriter := &SpyCountdownOperations{}

		CountDown(spySleepWriter, spySleepWriter)
		expectedOperations := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if spySleepWriter.Called != 3 {
			t.Errorf("Expected to be %d times, but was called %d times", 3, spySleepWriter.Called)
		}
		if !reflect.DeepEqual(expectedOperations, spySleepWriter.Operations) {
			t.Errorf("Expected operations to be %v, but was %v", expectedOperations, spySleepWriter.Operations)
		}
	})
}
