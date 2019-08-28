package cmd

import (
	"github.com/crerwin/distributedtranscoding/pkg/api"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "commands for the DTC API",
}

var apiServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the DTC API",
	Run:   apiServeRun,
}

func init() {
	apiCmd.AddCommand(apiServeCmd)
}

func apiServeRun(cmd *cobra.Command, args []string) {
	api.Serve()
}
