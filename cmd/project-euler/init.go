package euler

import (
  "github.com/spf13/cobra"
)

// AddSubCommands Adds this project's commands as the subcommad to the parent
func AddSubCommands(parent *cobra.Command) {
  setOneCmdFlags()
  parent.AddCommand(oneCmd)

  setTwoCmdFlags()
  parent.AddCommand(twoCmd)

  setTwentyFiveFlags()
  parent.AddCommand(twentyFiveCmd)
}
