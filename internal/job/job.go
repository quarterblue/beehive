package job

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
	Stop() error
}

type DockerJob struct {
	// Global unique ID to identify a docker job
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	State          State     `json:"state"`
	Owner          string    `json:"owner"`
	CreateTime     time.Time `json:"created_time"`
	LastStartTime  time.Time `json:"last_start_time"`
	LastFinishTime time.Time `json:"last_finish_time"`
}

type JobEvent struct {
	ID        int64
	State     State
	Timestamp time.Time
	DockerJob DockerJob
}

type Config struct {
	Name          string
	Cmd           []string
	Image         string
	Memory        int
	Disk          int
	Priority      int
	Container     string
	RestartPolicy string
	Arguments     string
}

type DockerContainer struct {
	Client      *client.Client
	Config      Config
	ContainerId string
}

func (d *DockerContainer) Run() Result {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(
		ctx, d.Config.Image, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Error pulling images")
		return Result{Error: err}
	}
	io.Copy(os.Stdout, reader)
	return Result{
		Error:  nil,
		Status: "Finished",
	}
}

type Result struct {
	Error  error
	Name   string
	Status string
}
