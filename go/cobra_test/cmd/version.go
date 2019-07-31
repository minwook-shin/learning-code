package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Long text...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version v0.1")
	},
}
