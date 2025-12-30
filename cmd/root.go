// Package cmd is the primary package that deals with the handling and registration of cobra commands.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cpt",
	Short: "cpt is a command line interface that bundles together cryptographic tools such as hashing, (de)ciphering, key generation, among others.",
	Long:  "cpt is a command line interface that bundles together cryptographic tools such as hashing, (de)ciphering, key generation, among others.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newHashCmd())
	rootCmd.AddCommand(newCipherCmd())
}
