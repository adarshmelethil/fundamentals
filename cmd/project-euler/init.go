package euler

import (
	"github.com/spf13/cobra"
)

func AddSubCommands(parent *cobra.Command) {
	setOneCmdFlags()
	parent.AddCommand(oneCmd)
}

