package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  RunE: func(cmd *cobra.Command, args []string) error {
    fmt.Println("1.0.0")
    return nil
  },
}