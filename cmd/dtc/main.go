package main

// CLI executable

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"
	"github.com/crerwin/distributedtranscoding/pkg/executors"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Infof("distributed transcoding cli version %v", dtc.Version)
	e := executors.NewVideoTranscodingExecutor()
	fmt.Print(e.DetectCrop())
}
