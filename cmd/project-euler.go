package cmd

import (
  "github.com/spf13/cobra"

  euler "github.com/adarshmelethil/fundamentals/cmd/project-euler"
)

var pEulerCmd = &cobra.Command{
  Use: "project-euler",
	Aliases: []string{"pe"},
  Short: "project euler problems",
  Run: DelegateSubcommands,
}

func init() {
  rootCmd.AddCommand(pEulerCmd)
  euler.AddSubCommands(pEulerCmd)
}
