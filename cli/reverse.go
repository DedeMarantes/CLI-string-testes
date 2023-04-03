package cli

import (
	"Go-Testes/functions"
	"fmt"

	"github.com/spf13/cobra"
)

var Reverse = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Inverte uma string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resultado := functions.ReverseString(args[0])
		fmt.Println(resultado)
	},
}
