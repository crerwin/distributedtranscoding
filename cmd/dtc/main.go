package main

// CLI executable

import (
	"log"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) // no timestamps on cli
	log.Printf("distributed transcoding cli version %v", dtc.Version)

}
