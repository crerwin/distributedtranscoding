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
	return ce.command + " mock output", nil
}
