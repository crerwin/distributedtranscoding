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

func kubernetesJobListRun(cmd *cobra.Command, args []string) {
	dtc.NewKubeClient().GetJobs()
}

func kubernetesInitRun(cmd *cobra.Command, args []string) {
	dtc.NewKubeClient().Init()
}

func kubernetesJobCreateRun(cmd *cobra.Command, args []string) {
	//c := dtc.NewKubeClient()
	i := new(dtc.Item)
	i.FileName = filepath.Base(args[0])
	i.Crop, _ = cmd.Flags().GetString("crop")
	i.Filters, _ = cmd.Flags().GetStringSlice("filters")
	i.ForcedRate, _ = cmd.Flags().GetString("forcedRate")
	j := new(dtc.Job)
	j.InboxPath = "inbox/"
	j.OutboxPath = "outbox/"
	j.ItemSubPath = "Frasier/s1d1/"
	j.Item = i
	//c.CreateTranscodeJob(i)
	fmt.Printf("file: %v\n", i.FileName)
	fmt.Printf("crop: %v\n", i.Crop)
	fmt.Printf("filters: %v\n", i.Filters)
	fmt.Printf("forcedRate: %v\n", i.ForcedRate)
	fmt.Printf("inputpath: %v\n", j.InboxPath+j.ItemSubPath+j.Item.FileName)
	fmt.Printf("outputpath: %v\n", j.OutboxPath+j.ItemSubPath+j.Item.FileName)
}
