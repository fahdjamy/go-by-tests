package main

import "fmt"

func main() {
	fmt.Println(RunnerWithPort(""))
}

const DefaultPort = "8090"
const RunningPrefix = "running..."

func RunnerWithPort(port string) string {
	if port == "" {
		port = DefaultPort
	}
	return RunningPrefix + " on port, :" + port
}
