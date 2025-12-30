package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCipherCmd() *cobra.Command {
	var (
		filepath  string
		algorithm string
		list      bool
	)

	cmd := &cobra.Command{
		Use:   "cipher",
		Short: "Temporary description",
		Long:  "Temporary description",
		// TODO: add description
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Cipher command ran")
			return nil
		},
	}

	cmd.Flags().StringVarP(&filepath, "file", "f", "", "Hash a specific file.")
	cmd.Flags().StringVarP(&algorithm, "algorithm", "a", "rot13", "Which cipher algorithm to use (e.g. rot13).")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "Lists all available algorithms.")

	return cmd
}
