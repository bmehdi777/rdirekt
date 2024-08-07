package client

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use: "rdirekt",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the number version",
		Run: func(cmd *cobra.Command, arg []string) {
			log.Println("1.0.0")
		},
	}

	var httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Proxyfy http request",
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
		Run: func(cmd *cobra.Command, arg []string) {
			log.Println("open port !")
			log.Println("args : ", arg)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(httpCmd)
	rootCmd.Execute()
}
