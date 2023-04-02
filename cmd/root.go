package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gonvert",
	Short: "Convert between JSON, YAML",
	Long:  "Convert between JSON, YAML",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		json2yamlCmd,
		yaml2jsonCmd,
	)
}
