package main

import "testing"

func TestRunnerMsg(t *testing.T) {
	received := RunnerMsg()
	expected := RunningPrefix

	if received != expected {
		t.Errorf("Recieved %q, Expected %q", received, expected)
	}
}

func TestRunningWithPort(t *testing.T) {
	got := RunnerWithPort(":8000")
	wanted := RunningPrefix + " on port, :8000"

	if got != wanted {
		t.Errorf("GOT %q, WANTED %q", got, wanted)
	}
}
