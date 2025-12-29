package cmd

import (
	"fmt"
	"strings"

	"github.com/TetieWasTaken/cpt/internal/hash"
	"github.com/spf13/cobra"
)

var algorithm string

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.ErrOrStderr(), "Attempting to hash using %q\n", algorithm)
		hasher, ok := hash.GetAlgorithm(algorithm)

		if !ok {
			fmt.Errorf("Unknown algorithm: %q", algorithm)
			return

			// TODO: Also return help menu or list of algorithms
		}

		// TODO: allow other types of input (e.g. file)

		data := strings.Join(args, " ")
		reader := strings.NewReader(data)

		sum, err := hasher.Hash(reader)
		if err != nil {
			println("error")
			println(err.Error())
			fmt.Println(err)
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), hash.HexDigest(sum))
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	hashCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "sha256", "Which hash algorithm to use (e.g. sha256)")
}
