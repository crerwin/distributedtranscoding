package executors

import (
	"log"
	"os/exec"
)

type cmdExecutor struct {
	command string
}

func NewCmdExecutor(command string) *cmdExecutor {
	ce := new(cmdExecutor)
	path, err := exec.LookPath(command)
	if err != nil {
		log.Printf("didn't find %s executable\n", command)
	} else {
		log.Printf("%s is executable and is in '%s'\n", command, path)
		ce.command = command
	}
	return ce
}
