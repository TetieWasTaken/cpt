package clio

import (
	"fmt"
	"io"
)

func Close(closer io.Closer, onErr func(error)) {
	if closer == nil {
		return
	}

	if err := closer.Close(); err != nil && onErr != nil {
		onErr(fmt.Errorf("error while trying to close: %w", err))
	}
}
