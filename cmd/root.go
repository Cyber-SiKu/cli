package cmd

import (
	"fmt"
	"os"

	list "github.com/Cyber-SiKu/cli/cmd/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	confPath     string
	format 		 string

	RootCmd = &cobra.Command{
		Use:   "curvefs_tool",
		Short: "A tool for curvefs",
		Long: `curvefs tool is a tool for curvefs.`,
	}
)

// Execute executes the root command.
func Execute() error {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&confPath, "conf", "", "config file (default is $HOME/.curve/curve.yaml)")
	RootCmd.PersistentFlags().StringVar(&format, "format", "plain", "output type (json,plain)")
	viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	addSubCommands()
	return RootCmd.Execute()
}

func initConfig() {
	if confPath != "" {
		// Use config file from the flag.
		fmt.Println(confPath)
		viper.SetConfigFile(confPath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".curvefs" (without extension).
		viper.AddConfigPath(home+"/.curve")
		viper.SetConfigType("yaml")
		viper.SetConfigName("curve")
	}

	viper.AutomaticEnv()
	viper.SetDefault("format", format)
	
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func addSubCommands() {
	RootCmd.AddCommand(list.NewListCmd())
}