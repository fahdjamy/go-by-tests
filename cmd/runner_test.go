package main

import "testing"

func TestRunnerMsg(t *testing.T) {
	t.Run("running with empty port should default to port 8090", func(t *testing.T) {
		received := RunnerWithPort("")
		expected := RunningPrefix + " on port, :8090"

		assertEqualStrings(t, received, expected)
	})

	t.Run("running with a port specified", func(t *testing.T) {
		received := RunnerWithPort("8000")
		expected := RunningPrefix + " on port, :8000"

		assertEqualStrings(t, received, expected)
	})
}

func assertEqualStrings(t testing.TB, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("Recieved %q, Expected %q", received, expected)
	}
}
