package cli

import (
	"fmt"
	"log/slog"

	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/fa"
	"github.com/polnaya-katuxa/bmstu/sem_02_mag/compilers/lab_01/internal/regexp"
	"github.com/spf13/cobra"
)

var Mode string

func init() {
	rootCmd.AddCommand(setRegexpCmd)
	setRegexpCmd.Flags().StringVarP(&Mode, "mode", "m", "", "Logs mode")
}

var setRegexpCmd = &cobra.Command{
	Use:   "regexp",
	Short: "Set regexp for pupipu",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("need regexp to set: %w", err)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		in := cmd.Flags().Arg(0)
		regexp := &regexp.Regexp{
			Initial: "(" + in + ")" + "#",
		}

		slog.Info("start building DFA by regexp", slog.String("regexp", in))

		err := fa.BuildFaByRegexp(regexp, Mode)
		if err != nil {
			fmt.Println(Fail + "invalid regexp" + EndC)
		} else {
			fmt.Println(OK + fmt.Sprintf("set new regexp: %s", in) + EndC)
		}

		return nil
	},
}
