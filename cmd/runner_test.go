package main

import (
	"fmt"
	"testing"
)

func TestRunnerMsg(t *testing.T) {
	t.Run("running with empty port should default to port 8090", func(t *testing.T) {
		received := RunnerWithPort("", "")
		expected := RunningPrefix + " on port, :8090"

		assertEqualStrings(t, received, expected)
	})

	t.Run("running with a port specified", func(t *testing.T) {
		received := RunnerWithPort("8000", "")
		expected := RunningPrefix + " on port, :8000"

		assertEqualStrings(t, received, expected)
	})

	t.Run("summing with valid values", func(t *testing.T) {
		response := Summer([]int{1, 2, 3, 4, 5})
		expected := 15

		if response != expected {
			t.Errorf("Summer did not return want value. Expected: %d, Actual: %d", expected, response)
		}
	})
}

func assertEqualStrings(t testing.TB, received, expected string) {
	// this helper is need to tell the test suite that this method is a helper
	// By doing this when it fails the line number reported will be in our function call rather than
	// inside our test helper
	t.Helper()
	if received != expected {
		t.Errorf("Recieved %q, Expected %q", received, expected)
	}
}

// note that the example function will not be executed if you remove the comment
// Output: running... on port, :1080
func ExampleRunnerWithPort() {
	outPutRunningStr := RunnerWithPort("1080", "")
	fmt.Println(outPutRunningStr)
	// Output: running... on port, :1080
}
