package mock

type cmdExecutor struct {
	command string
}

func NewCmdExecutor(command string) *cmdExecutor {
	ce := new(cmdExecutor)
	ce.command = command
	return ce
}

func (ce *cmdExecutor) Execute(args ...string) (string, error) {
	cmd := ce.command
	for _, a := range args {
		cmd += " " + a
	}
	return cmd, nil
}
