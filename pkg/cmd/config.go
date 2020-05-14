package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Displays config values",
	Long:  "Displays config values.",
	Run:   configRun,
}

func configRun(cmd *cobra.Command, args []string) {
	fmt.Println(viper.AllSettings())
}
