package cmd

import (
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
	c := dtc.NewKubeClient()
	i := new(dtc.Item)
	i.InputFile = "/data/inbox/frasier/s1d1/Frasier - s01e01.mkv"
	i.OutputFile = "/data/outbox/frasier/s1d1/Frasier - s01e01.mkv"
	i.Crop = "0:0:0:0"
	c.CreateTranscodeJob(i)
}
