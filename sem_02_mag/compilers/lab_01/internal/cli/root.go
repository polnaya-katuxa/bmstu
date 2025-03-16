package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pupipu",
	Short: "Pupipu helps to build DFA by regexp and parse strings.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nhello, i will parse your regexps!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
