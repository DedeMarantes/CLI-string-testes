package cli

import "github.com/spf13/cobra"

var str_util = &cobra.Command{
	Use:   "strutil",
	Short: "Utilities for strings",
}

var sort = Sort()

func Execute() error {
	return str_util.Execute()
}

func init() {
	cobra.OnInitialize()
	str_util.AddCommand(Reverse)
	str_util.AddCommand(sort)
}
