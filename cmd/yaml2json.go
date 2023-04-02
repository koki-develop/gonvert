package cmd

import (
	"fmt"

	"github.com/koki-develop/gonvert/internal/json"
	"github.com/spf13/cobra"
)

var yaml2jsonCmd = &cobra.Command{
	Use:          "yaml2json",
	Short:        "Convert YAML to JSON",
	Long:         "Convert YAML to JSON.",
	Aliases:      []string{"yamltojson", "yml2json", "ymltojson", "y2j", "ytoj", "yj"},
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := open(args)
		if err != nil {
			return err
		}

		j, err := json.New(&json.JSONConfig{Indent: 2, Minify: false}).FromYAML(r)
		if err != nil {
			return err
		}

		fmt.Print(string(j))
		return nil
	},
}
