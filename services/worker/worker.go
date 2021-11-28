package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/quarterblue/beehive/internal/job"

	"github.com/quarterblue/beehive/services/worker/pb"
)

type Worker struct {
	ID        string
	Name      string
	Map       map[uuid.UUID]string
	TaskCount int
	mu        sync.Mutex
}

func NewWorker(id, name string) *Worker {
	return &Worker{
		ID:        id,
		Name:      name,
		Map:       make(map[uuid.UUID]string),
		TaskCount: 0,
		mu:        sync.Mutex{},
	}
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

func (w *Worker) MachineSpec(ctx context.Context, request *pb.SpecRequest) (*pb.SpecResponse, error) {
	mSpec, err := retrieveSpec()
	if err != nil {
		return nil, err
	}

	return mSpec, nil
}
