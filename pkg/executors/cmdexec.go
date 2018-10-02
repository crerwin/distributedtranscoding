package executors

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type cmdExecutor struct {
	command string
}

func NewCmdExecutor(command string) *cmdExecutor {
	ce := new(cmdExecutor)
	path, err := exec.LookPath(command)
	if err != nil {
		log.Warnf("didn't find %s executable\n", command)
	} else {
		log.Debugf("%s is executable and is in '%s'\n", command, path)
		ce.command = command
	}
	return ce
}

func (ce *cmdExecutor) Execute(args ...string) (string, error) {
	cmd := exec.Command(ce.command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Warnf("cmd.Run() failed with %s\n", err)
	}
	return strings.TrimSuffix(string(out), "\n"), err
}
