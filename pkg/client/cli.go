package client

import (
	"github.com/spf13/cobra"
)

func StartCmd()  {
	var rootCmd = &cobra.Command {
		Use: "rdirekt-runner",
	}

	rootCmd.Execute()
}
