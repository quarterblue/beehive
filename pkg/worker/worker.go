package worker

import (
	"fmt"

	"github.com/google/uuid"
)

type Worker struct {
	Map       map[uuid.UUID]string
	TaskCount int
}

func (w *Worker) RunJob() {
	fmt.Println("Run Job!")
}

func (w *Worker) StartJob() {
	fmt.Println("Start Job!")
}

func (w *Worker) StopJob() {
	fmt.Println("Stop Job!")
}
