package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/koki-develop/gonvert/internal/json"
	"github.com/spf13/cobra"
)

var yaml2jsonCmd = &cobra.Command{
	Use:     "yaml2json",
	Aliases: []string{"yml2json", "y2j"},
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

		j, err := json.FromYAML(r)
		if err != nil {
			return err
		}

		fmt.Print(string(j))
		return nil
	},
}
