package main

import (
	"fmt"
	"io"
)

func Communicator(writer io.Writer, msg string) error {
	_, err := fmt.Fprintf(writer, "Incoming message: %s", msg)
	return err
}
