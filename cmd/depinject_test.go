package main

import (
	"bytes"
	"testing"
)

func TestCommunicator(t *testing.T) {
	t.Run("communicator writing", func(t *testing.T) {
		buf := bytes.Buffer{}
		msg := "Go home"
		err := Communicator(&buf, msg)

		assertNoError(t, err)

		written := buf.String()
		expected := "Incoming message: " + msg
		if written != expected {
			t.Errorf("expected %q, got %q", expected, written)
		}
	})
}
