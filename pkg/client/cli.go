package client

import (
	"errors"
	"log"
	"strconv"

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
		Use:   "http <PORT>",
		Short: "Proxyfy http request",
		Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1), func(cmd *cobra.Command, arg []string) error {
			if value, err := strconv.Atoi(arg[0]); err != nil {
				return errors.New("PORT should be a numeric value")
			} else if value <= 0 || value > 65535 {
				return errors.New("PORT should be between 0 and 65535")
			}
			return nil
		}),
		Run: func(cmd *cobra.Command, arg []string) {
			log.Println("open port !")
			log.Println("args : ", arg)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(httpCmd)
	rootCmd.Execute()
}
