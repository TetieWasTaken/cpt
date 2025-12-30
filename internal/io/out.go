package io

import (
	"io"
	"os"
)

func OpenOutput(path string, out io.Writer) (io.WriteCloser, error) {
	if path == "" {
		return nopWriteCloser{out}, nil
	}

	return os.Create(path)
}

type nopWriteCloser struct{ io.Writer }

// stdout cannot be closed, hence an empty function to match type io.WriteCloser.
func (n nopWriteCloser) Close() error { return nil }
