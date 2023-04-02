package cmd

import (
	"fmt"
	"io"
	"os"

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
		var r io.Reader
		if len(args) == 0 {
			r = os.Stdin
		} else {
			f, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer f.Close()
			r = f
		}

		j, err := json.New(&json.JSONConfig{Indent: 2, Minify: false}).FromYAML(r)
		if err != nil {
			return err
		}

		fmt.Print(string(j))
		return nil
	},
}
