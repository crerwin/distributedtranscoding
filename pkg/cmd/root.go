package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dtc",
	Short: "DTC is a tool for distributing transcoding jobs across a set of compute resources.",
	Long:  "DTC is a tool that uses Don Melton's Video Transcoding tools to transcode videos.  It coordinates transcode jobs across multiple compute resources.",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("hello")
	// },
}

func init() {
	rootCmd.AddCommand(versionCmd)
	redisCmd.AddCommand(redisPingCmd)
	redisCmd.AddCommand(redisInitCmd)
	redisCmd.AddCommand(redisAddCmd)
	rootCmd.AddCommand(redisCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
