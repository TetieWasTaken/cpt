package cmd

import (
	"github.com/TetieWasTaken/cpt/internal/logger"
	"github.com/spf13/cobra"
)

func vlogf(cmd *cobra.Command) func(string, ...any) error {
	return func(format string, a ...any) error {
		return logger.Vprintf(cmd, format, a...)
	}
}
