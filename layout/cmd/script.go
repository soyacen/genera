package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use: "script",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("script called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scriptCmd)
}
