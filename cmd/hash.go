package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TetieWasTaken/cpt/internal/hash"
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
				fmt.Println(cmd.OutOrStdout(), strings.Join(hash.ListAlgorithms(), "\n"))
				return nil
			}

			hasher, ok := hash.GetAlgorithm(algorithm)

			if !ok {
				return fmt.Errorf("unknown algorithm: %q", algorithm)
			}

			var reader io.Reader

			if filepath != "" {
				file, err := os.Open(filepath)
				if err != nil {
					return err
				}

				defer func() {
					if err := file.Close(); err != nil {
						fmt.Println(cmd.ErrOrStderr(), err)
					}
				}()

				reader = file
			} else if len(args) > 0 {
				reader = strings.NewReader(strings.Join(args, " "))
			} else if stat, err := os.Stdin.Stat(); err == nil && stat.Mode()&os.ModeCharDevice == 0 {
				reader = bufio.NewReader(os.Stdin)
			} else {
				return fmt.Errorf("no input found, please provide data to hash")
			}

			sum, err := hasher.Hash(reader)
			if err != nil {
				return err
			}

			writer := cmd.OutOrStdout()
			if out != "" {
				f, err := os.Create(out)
				if err != nil {
					return err
				}

				defer f.Close()
				writer = f

				fmt.Fprintf(cmd.ErrOrStderr(), "Wrote output to %s\n", out)
			}

			_, err = fmt.Fprintln(writer, hash.HexDigest(sum))

			return err
		},
	}

	cmd.Flags().StringVarP(&filepath, "file", "f", "", "Hash the contents of a specific file.")
	cmd.Flags().StringVarP(&algorithm, "algorithm", "a", "sha256", "Which hash algorithm to use (e.g. sha256).")
	cmd.RegisterFlagCompletionFunc(
		"algorithm",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return hash.ListAlgorithms(), cobra.ShellCompDirectiveNoFileComp
		},
	)
	cmd.Flags().BoolVarP(&list, "list", "l", false, "Lists all available algorithms.")
	cmd.Flags().StringVarP(&out, "out", "o", "", "Output to a specific file.")
	// TODO: add --verbose

	return cmd
}
