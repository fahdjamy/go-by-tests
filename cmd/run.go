package main

import "fmt"

func main() {
	fmt.Println(RunnerMsg())
}

const RunningPrefix = "running..."

func RunnerMsg() string {
	return RunningPrefix
}

func RunnerWithPort(port string) string {
	return RunningPrefix + " on port, " + port
}
