package baseClass

import (
	"github.com/spf13/cobra"
)

type CmdFunctions interface {
	Init(cmd *cobra.Command, args []string) error
	RunCommand(cmd *cobra.Command, args []string) error
	Print(cmd *cobra.Command, args []string) error
}

type CmdValue struct {
	Use string
	Short string
	Long string
	Output []byte // Output of the command, type is json
	Format string
	Cmd *cobra.Command
}

func NewCmd(cv *CmdValue, cf CmdFunctions) *cobra.Command {
	return &cobra.Command{
		Use:   cv.Use,
		Short: cv.Short,
		Long:  cv.Long,
		PreRunE: cf.Init,
		RunE: cf.RunCommand,
		PostRunE: cf.Print,
		SilenceUsage : false,
		SilenceErrors: false,
		DisableFlagsInUseLine: true,
	}
}
