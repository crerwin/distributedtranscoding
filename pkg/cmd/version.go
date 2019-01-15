package cmd

import (
	"fmt"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the DTC version",
	Long:  "Displays the DTC version",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("DTC version %v\n", dtc.Version)
}
