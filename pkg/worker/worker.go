package worker

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/quarterblue/beehive/pkg/job"
)

type Worker struct {
	ID        string
	Name      string
	Map       map[uuid.UUID]string
	TaskCount int
}

func (w *Worker) RunJob() {
	fmt.Println("Run Job!")
}

func (w *Worker) StartJob(j job.Job) job.Result {
	fmt.Println("Start Job!")
	return job.Result{
		Error:  nil,
		Name:   "Done",
		Status: "Done",
	}
}

func (w *Worker) StopJob() {
	fmt.Println("Stop Job!")
}
