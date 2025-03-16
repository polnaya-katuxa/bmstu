package cli

import (
	"fmt"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/fa"
	"github.com/spf13/cobra"
)

const (
	OK   = "\033[102m"
	Fail = "\033[101m"
	EndC = "\033[0m"
)

func init() {
	rootCmd.AddCommand(processSeqCmd)
}

var processSeqCmd = &cobra.Command{
	Use:   "process",
	Short: "Process sequence",
	Long:  `Use 'pupipu regexp <regexp>' before processing`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("need sequence to process: %w", err)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		dfa, err := fa.Read()
		if err != nil {
			return fmt.Errorf("read dfa: %w", err)
		}

		ok := dfa.Model(cmd.Flags().Arg(0))
		if ok {
			fmt.Println(OK + "OK" + EndC)
		} else {
			fmt.Println(Fail + "FAIL" + EndC)
		}

		return nil
	},
}
