package cmd

import (
	"fmt"
	"os"

	"github.com/CrowdStrike/gqltools/cmd/compare"
	"github.com/CrowdStrike/gqltools/cmd/linter"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gql",
		Short: "gql is a CLI built for federated GraphQL services' schemas",
		Long:  `gql is a CLI built for federated GraphQL services' schemas`,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}
	cmd.AddCommand(linter.NewLintCmd())
	cmd.AddCommand(compare.NewCompareCmd())
	return cmd
}

// Execute is a wrapper to execute Run function in subcommands
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
