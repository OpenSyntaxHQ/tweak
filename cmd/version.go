package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of tweak",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tweak version %s\n", Version)
	},
}
