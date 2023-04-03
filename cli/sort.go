package cli

import (
	"Go-Testes/functions"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type SortOptions struct {
	InverseOrder bool
}

func Sort() *cobra.Command {
	opt := SortOptions{}
	var sort = &cobra.Command{
		Use:   "sort",
		Short: "sort a list of strings in alphabetical order",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			res := functions.SortStringList(args, opt.InverseOrder)
			fmt.Println(strings.Join(res, "\n")) //transforma lista em strings e quebra linhas
		},
	}
	sort.Flags().BoolVarP(&opt.InverseOrder, "reverse", "r", false, "Colocar em ordem reversa alfabética")
	return sort
}
