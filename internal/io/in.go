package clio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ParseInput processes user input from the CLI and returns a stream.
func ParseInput(args []string, file string) (io.ReadCloser, error) {
	if file != "" {
		// Open the specified file to read.
		return os.Open(file)
	}

	if len(args) > 0 {
		// Combine args into a single string and return as reader.
		return io.NopCloser(strings.NewReader(strings.Join(args, " "))), nil
	}

	stat, err := os.Stdin.Stat()
	// Check if stdin is not a terminal (e.g. via piping).
	if err == nil && stat.Mode()&os.ModeCharDevice == 0 {
		return io.NopCloser(bufio.NewReader(os.Stdin)), nil
	}

	return nil, fmt.Errorf("no input provided")
}
