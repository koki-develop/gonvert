package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/koki-develop/gonvert/internal/yaml"
	"github.com/spf13/cobra"
)

var json2yamlCmd = &cobra.Command{
	Use:     "json2yaml",
	Aliases: []string{"json2yml", "j2y"},
	Args:    cobra.MaximumNArgs(1),
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

		y, err := yaml.FromJSON(r)
		if err != nil {
			return err
		}
		fmt.Print(string(y))
		return nil
	},
}
