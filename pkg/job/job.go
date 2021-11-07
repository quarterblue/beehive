package job

import (
	"time"

	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)

type Job interface {
	Execute() error
}

type DockerJob struct {
	// Global unique ID to identify a docker job
	ID uuid.UUID

	Name string

	State State

	Image string

	Owner string

	Memory int

	Disk int

	Priority int

	Container string

	RestartPolicy string

	Arguments string

	CreateTime time.Time

	StartTime time.Time

	FinishTime time.Time
}
