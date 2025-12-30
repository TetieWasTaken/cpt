package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/TetieWasTaken/cpt/internal/hash"
	clio "github.com/TetieWasTaken/cpt/internal/io"
	logger "github.com/TetieWasTaken/cpt/internal/logger"
	"github.com/spf13/cobra"
)

func newHashCmd() *cobra.Command {
	var (
		filepath  string
		algorithm string
		list      bool
		out       string
	)

	cmd := &cobra.Command{
		Use:   "hash",
		Short: "Maps a string to a unique string with a fixed length that cannot be reversed.",
		Long: `Uses one-way deterministic algorithms to create a fixed-length string that:
			1. Cannot be feasibly reversed.
			2. Is unique to the input.

			One common application of cryptographic hash functions is to store passwords safely.

			Example:
			cpt hash "The quick brown fox jumps over the lazy dog"`,
		Args: func(cmd *cobra.Command, args []string) error {
			if filepath == "" && len(args) == 0 {
				if stat, err := os.Stdin.Stat(); err == nil && stat.Mode()&os.ModeCharDevice == 0 {
					return nil
				}

				return fmt.Errorf("no input provided")
			}

			if filepath != "" && len(args) > 0 {
				return fmt.Errorf("cannot use both file and arguments, please only enter one input source")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if list {
				fmt.Fprintln(cmd.OutOrStdout(), strings.Join(hash.ListAlgorithms(), "\n"))
				return nil
			}

			hasher, ok := hash.GetAlgorithm(algorithm)

			if !ok {
				return fmt.Errorf("unknown algorithm: %q", algorithm)
			}

			if err := logger.Vprintf(cmd, "Using algorithm %s", algorithm); err != nil {
				return err
			}

			reader, err := clio.ParseInput(args, filepath, func(format string, a ...any) error {
				return logger.Vprintf(cmd, format, a...)
			})
			if err != nil {
				return err
			}

			defer clio.Close(reader, func(err error) {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
			})

			sum, err := hasher.Hash(reader)
			if err != nil {
				return err
			}

			writer, err := clio.OpenOutput(out, cmd.OutOrStdout(), func(format string, a ...any) error {
				return logger.Vprintf(cmd, format, a...)
			})
			if err != nil {
				return err
			}

			defer clio.Close(writer, func(err error) {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
			})

			_, err = fmt.Fprintln(writer, hash.HexDigest(sum))

			return err
		},
	}

	cmd.Flags().StringVarP(&filepath, "file", "f", "", "Hash the contents of a specific file.")
	cmd.Flags().StringVarP(&algorithm, "algorithm", "a", "sha256", "Which hash algorithm to use (e.g. sha256).")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "Lists all available algorithms.")
	cmd.Flags().StringVarP(&out, "out", "o", "", "Output to a specific file.")

	return cmd
}
