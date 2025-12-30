package clio

import (
	"io"
	"os"
)

func OpenOutput(path string, out io.Writer, vlog func(format string, a ...any) error) (io.WriteCloser, error) {
	if path == "" {
		return nopWriteCloser{out}, nil
	}

	if vlog != nil {
		_ = vlog("Writing output to %s.", path)
	}

	return os.Create(path)
}

type nopWriteCloser struct{ io.Writer }

// stdout cannot be closed, hence an empty function to match type io.WriteCloser.
func (n nopWriteCloser) Close() error { return nil }
