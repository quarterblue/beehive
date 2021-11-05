package cmd

type Job interface {
	Execute() error
}

type DockerJob struct {
	Name string

	Owner string

	Container string

	Image string

	Arguments string

	Id string

	Status string

	Datetime string

	Result string
}
