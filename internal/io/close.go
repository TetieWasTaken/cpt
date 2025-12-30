package clio

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func Close(closer io.Closer, onErr func(error)) {
	if closer == nil {
		return
	}

	if err := closer.Close(); err != nil && onErr != nil {
		onErr(fmt.Errorf("error while trying to close: %w", err))
	}
}

func CloseWithDefer(cmd *cobra.Command, closer io.Closer) func() {
	return func() {
		Close(closer, func(err error) { fmt.Fprintln(cmd.ErrOrStderr(), err) })
	}
}
