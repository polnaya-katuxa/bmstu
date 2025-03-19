package cli

import (
	"fmt"
	"log/slog"

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
		slog.Info("reading DFA from temp/dfa.json")

		dfa, err := fa.Read()
		if err != nil {
			return fmt.Errorf("read dfa: %w", err)
		}

		in := cmd.Flags().Arg(0)

		slog.Info("modeling DFA on sequence", slog.String("input", in))

		ok := dfa.Model(in)
		fmt.Print("\nANSWER: ")
		if ok {
			fmt.Println(OK + "OK" + EndC)
		} else {
			fmt.Println(Fail + "FAIL" + EndC)
		}

		return nil
	},
}
