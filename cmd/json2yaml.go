package cmd

import (
	"github.com/spf13/cobra"
)

var json2yamlCmd = &cobra.Command{
	Use:     "json2yaml",
	Aliases: []string{"json2yml", "j2y"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
