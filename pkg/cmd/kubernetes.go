package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"
	"github.com/spf13/cobra"
)

var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Interact with Kubernetes",
}

var kubernetesJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Interact with Kubernetes jobs",
}

var kubernetesJobListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of jobsjobs",
	Run:   kubernetesJobListRun,
}

var kubernetesJobCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a job",
	Run:   kubernetesJobCreateRun,
}

var kubernetesInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize",
	Run:   kubernetesInitRun,
}

func init() {
	kubernetesCmd.AddCommand(kubernetesInitCmd)
	kubernetesCmd.AddCommand(kubernetesJobCmd)
	kubernetesJobCmd.AddCommand(kubernetesJobListCmd)
	kubernetesJobCmd.AddCommand(kubernetesJobCreateCmd)
	kubernetesJobCreateCmd.Flags().StringP("crop", "c", "0:0:0:0",
		"Crop n:n:n:n")
	kubernetesJobCreateCmd.Flags().StringSliceP("filters", "f", []string{},
		"Filters")
	kubernetesJobCreateCmd.Flags().StringP("forcedRate", "r", "",
		"Framerate to force")
}

func kubernetesJobListRun(cmd *cobra.Command, args []string) {
	dtc.NewKubeClient().GetJobs()
}

func kubernetesInitRun(cmd *cobra.Command, args []string) {
	dtc.NewKubeClient().Init()
}

func kubernetesJobCreateRun(cmd *cobra.Command, args []string) {
	workspacePath, _ := cmd.Flags().GetString("workspace")
	c := dtc.NewKubeClient()
	j := new(dtc.Job)
	j.InboxPath = "inbox/"
	j.OutboxPath = "outbox/"
	j.Item = dtc.NewItemFromPath(args[0], filepath.Join(workspacePath, j.InboxPath))
	j.Item.Crop, _ = cmd.Flags().GetString("crop")
	j.Item.Filters, _ = cmd.Flags().GetStringSlice("filters")
	j.Item.ForcedRate, _ = cmd.Flags().GetString("forcedRate")

	fmt.Printf("file: %v\n", j.Item.FileName)
	fmt.Printf("crop: %v\n", j.Item.Crop)
	fmt.Printf("filters: %v\n", j.Item.Filters)
	fmt.Printf("forcedRate: %v\n",
		j.Item.ForcedRate)
	fmt.Printf("inputpath: %v\n",
		filepath.Join(workspacePath, j.InboxPath, j.Item.SubPath,
			j.Item.FileName))
	fmt.Printf("outputpath: %v\n",
		filepath.Join(workspacePath, j.OutboxPath, j.Item.SubPath,
			j.Item.FileName))
	c.CreateTranscodeJob(j)
}
