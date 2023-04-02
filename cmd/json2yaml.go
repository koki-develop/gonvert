package cmd

import (
	"fmt"

	"github.com/koki-develop/gonvert/internal/yaml"
	"github.com/spf13/cobra"
)

var json2yamlCmd = &cobra.Command{
	Use:          "json2yaml",
	Short:        "Convert JSON to YAML",
	Long:         "Convert JSON to YAML.",
	Aliases:      []string{"jsontoyaml", "json2yml", "jsontoyml", "j2y", "jtoy", "jy"},
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := open(args)
		if err != nil {
			return err
		}

		y, err := yaml.New(&yaml.YAMLConfig{Indent: 2}).FromJSON(r)
		if err != nil {
			return err
		}
		fmt.Print(string(y))
		return nil
	},
}
