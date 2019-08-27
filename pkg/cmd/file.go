package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Process files",
}

var fileOutboxPrepCmd = &cobra.Command{
	Use:   "outboxprep",
	Short: "Create the file's subpath in the outbox to prepare for it to be transcoded",
	Run:   fileOutboxPrepRun,
}

var fileDispatchCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Directly dispatch a file to a worker to be transcoded",
}

func init() {
	fileCmd.AddCommand(fileOutboxPrepCmd)
}

func fileOutboxPrepRun(cmd *cobra.Command, args []string) {
	workspacePath, _ := cmd.Flags().GetString("workspace")
	inboxPath := "inbox/"
	outboxPath := "outbox/"
	fmt.Println(filepath.Join(workspacePath, inboxPath))
	i := dtc.NewItemFromPath(args[0], filepath.Join(workspacePath, inboxPath))
	newDir := filepath.Join(workspacePath, outboxPath, i.SubPath)
	fmt.Printf("Making directory: %v\n", newDir)
	os.MkdirAll(newDir, 0700)
}
