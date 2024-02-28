package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"
	luganda = "Luganda"

	spanishPrefix = "correr..."
	englishPrefix = "running..."
	lugandaPrefix = "edukira..."
	frenchPrefix  = "en cours d'exécution..."
)

func main() {
	fmt.Println(RunnerWithPort("", ""))
}

const DefaultPort = "8090"
const RunningPrefix = "running..."

func RunnerWithPort(port, language string) string {
	if port == "" {
		port = DefaultPort
	}
	runningOn := runningOnPrefix(language)
	return runningOn + " on port, :" + port
}

func runningOnPrefix(language string) string {
	runningOn := RunningPrefix
	switch language {
	case french:
		runningOn = frenchPrefix
	case spanish:
		runningOn = spanishPrefix
	case luganda:
		runningOn = lugandaPrefix
	default:
		runningOn = englishPrefix
	}
	return runningOn
}
