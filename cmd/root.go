package cmd

import (
	"fmt"
	"os"

	"github.com/OpenSyntaxHQ/tweak/ui"
	"github.com/spf13/cobra"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "tweak",
	Short: "tweak is a fast and flexible string/text transformer",
	Long: `tweak is a command line tool that allows you to quickly apply various
transformation operations on the input text.

Run without arguments for an interactive TUI, or use a subcommand directly:

  tweak md5 "Hello World"
  echo "Hello World" | tweak base64-encode
  tweak upper file.txt

Complete documentation is available at https://github.com/OpenSyntaxHQ/tweak`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			x := ui.New("")
			x.Render()
		}
		return nil
	},
}

func init() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
