package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/quarterblue/beehive/internal/job"
)

func (app *application) createJobHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name           string    `json:"name"`
		State          job.State `json:"state"`
		Owner          string    `json:"owner"`
		CreateTime     time.Time `json:"created_time"`
		LastStartTime  time.Time `json:"last_start_time"`
		LastFinishTime time.Time `json:"last_finish_time"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		log.Println(err)
		return
	}

	job := &job.DockerJob{
		Name:           input.Name,
		State:          input.State,
		Owner:          input.Owner,
		CreateTime:     input.CreateTime,
		LastStartTime:  input.LastStartTime,
		LastFinishTime: input.LastFinishTime,
	}

	err = app.models.Jobs.Insert(job)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/jobs/%d", job.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"job": job}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showJobHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	job, err := app.models.Jobs.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"job": job}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) startJobHandler(w http.ResponseWriter, r *http.Request) {
}

type JobModel struct {
	DB *sql.DB
}

func (j JobModel) Insert(job *job.DockerJob) error {

	query := `
		INSERT INTO jobs (job_name, state, owner, created_time, last_start_time, last_end_time)
		VALUES ($1, $2, $3, $4)
		RETURNING node_id`

	args := []interface{}{job.Name, job.State, job.Owner, job.CreateTime, job.LastStartTime, job.LastFinishTime}

	return j.DB.QueryRow(query, args...).Scan(&job.ID)
}

func (j JobModel) Get(id int64) (*job.DockerJob, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT job_id, job_name, state, created_time, last_start_time, last_end_time 
        FROM jobs 
        WHERE job_id = $1`

	var job job.DockerJob

	err := j.DB.QueryRow(query, id).Scan(
		&job.ID,
		&job.Name,
		&job.Owner,
		&job.CreateTime,
		&job.LastStartTime,
		&job.LastFinishTime,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &job, nil
}

func (j JobModel) Update(job *job.DockerJob) error {
	return nil
}

func (j JobModel) Delete(id uint64) error {
	return nil
}
