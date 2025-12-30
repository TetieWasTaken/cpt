package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/TetieWasTaken/cpt/internal/hash"
	clio "github.com/TetieWasTaken/cpt/internal/io"
	logger "github.com/TetieWasTaken/cpt/internal/logger"
	"github.com/spf13/cobra"
)

type hashOptions struct {
	filepath  string
	algorithm string
	list      bool
	outPath   string
	format    string
}

func newHashCmd() *cobra.Command {
	var flags hashOptions
	allowedFormats := []string{"hex", "base64"}

	cmd := &cobra.Command{
		Use:   "hash [input]",
		Short: "Computes a cryptographic hash.",
		Args: func(cmd *cobra.Command, args []string) error {
			if flags.filepath != "" && len(args) > 0 {
				return fmt.Errorf("cannot use both file and arguments, please only enter one input source")
			}

			if !slices.Contains(allowedFormats, flags.format) {
				return fmt.Errorf("format '%s' is not supported", flags.format)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if flags.list {
				fmt.Fprintln(cmd.OutOrStdout(), strings.Join(hash.ListAlgorithms(), "\n"))
				return nil
			}

			hasher, ok := hash.GetAlgorithm(flags.algorithm)

			if !ok {
				return fmt.Errorf("unknown algorithm: %q", flags.algorithm)
			}

			if err := logger.Vprintf(cmd, "Using algorithm %s", flags.algorithm); err != nil {
				return err
			}

			reader, err := clio.ParseInput(args, flags.filepath, vlogf(cmd))
			if err != nil {
				return err
			}

			defer clio.CloseWithDefer(cmd, reader)()

			sum, err := hasher.Hash(reader)
			if err != nil {
				return err
			}

			writer, err := clio.OpenOutput(flags.outPath, cmd.OutOrStdout(), vlogf(cmd))
			if err != nil {
				return err
			}

			defer clio.CloseWithDefer(cmd, writer)()

			var res string

			switch flags.format {
			case "hex":
				res = hash.HexDigest(sum)
			case "base64":
				res = hash.Base64Digest(sum)
			}

			_, err = fmt.Fprintln(writer, res)

			return err
		},
	}

	cmd.Flags().StringVarP(&flags.filepath, "input", "i", "", "Hash the contents of a specific file.")
	cmd.Flags().StringVarP(&flags.algorithm, "algorithm", "a", "sha256", "Which hash algorithm to use (e.g. sha256).")
	cmd.Flags().BoolVarP(&flags.list, "list", "l", false, "Lists all available algorithms.")
	cmd.Flags().StringVarP(&flags.outPath, "out", "o", "", "Output to a specific file.")
	cmd.Flags().StringVarP(&flags.format, "format", "f", "hex", "Format of the output.")

	return cmd
}
