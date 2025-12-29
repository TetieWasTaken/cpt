package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hash called")
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
}
