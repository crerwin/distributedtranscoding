package executors

type Executor interface {
	Execute(args ...string) (string, error)
}
