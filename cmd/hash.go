package cmd

import (
	"fmt"
	"strings"

	"github.com/TetieWasTaken/cpt/internal/hash"
	"github.com/spf13/cobra"
)

var (
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
			return fmt.Errorf("Unknown algorithm: %q", algorithm)

			// TODO: Also return help menu or list of algorithms
		}

		// TODO: allow other types of input (e.g. file)

		data := strings.Join(args, " ")
		reader := strings.NewReader(data)

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

	hashCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "sha256", "Which hash algorithm to use (e.g. sha256).")
	hashCmd.Flags().BoolVarP(&list, "list", "l", false, "Lists all available algorithms.")
}
