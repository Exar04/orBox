package cmd

import (

	"github.com/spf13/cobra"
    "github.com/Exar04/orBox/cmd/expcommand"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Touch typing",
		Long:  "play and test your typing speed",
	}
)

func Execute() error {
    rootCmd.AddCommand(expcommand.Expcommand)
	return rootCmd.Execute()
}


