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

var (
	filepath  string
	algorithm string
	list      bool
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Maps a string to a unique string with a fixed length that cannot be reversed.",
	Long: `Uses one-way deterministic algorithms to create a fixed-length string that:
		1. Cannot be feasibly reversed.
		2. Is unique to the input.

		One common application of cryptographic hash functions is to store passwords safely.

		Example:
		cpt hash "The quick brown fox jumps over the lazy dog"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if list {
			fmt.Println(hash.ListAlgorithms())
			return nil
		}

		fmt.Fprintf(cmd.ErrOrStderr(), "Attempting to hash using %q\n", algorithm)
		hasher, ok := hash.GetAlgorithm(algorithm)

		if !ok {
			return fmt.Errorf("unknown algorithm: %q", algorithm)
		}

		data := ""
		var reader io.Reader

		if filepath != "" {
			file, err := os.Open(filepath)
			if err != nil {
				return err
			}

			defer file.Close()

			reader = file
		} else if len(args) > 0 {
			data = strings.Join(args, " ")
			reader = strings.NewReader(data)
		} else if file, error := os.Stdin.Stat(); error == nil && file.Mode()&os.ModeCharDevice == 0 {
			reader = bufio.NewReader(os.Stdin)
		} else {
			return fmt.Errorf("no input found, please provide data to hash")
		}

		sum, err := hasher.Hash(reader)
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), hash.HexDigest(sum))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&filepath, "file", "f", "", "Hash a specific file.")
	hashCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "sha256", "Which hash algorithm to use (e.g. sha256).")
	hashCmd.Flags().BoolVarP(&list, "list", "l", false, "Lists all available algorithms.")
}
