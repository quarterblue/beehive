package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/quarterblue/beehive/internal/node"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

func (app *application) createNodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) editNodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) showNodeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	node := node.Node{
		ID:       fmt.Sprint(id),
		Name:     "practice",
		IpAddr:   "127.39.10.2",
		Port:     "3001",
		JobCount: 0,
		Spec:     nil,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"node": node}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "There was a problem with your request.", http.StatusInternalServerError)
	}
}

func (app *application) deleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

type NodeModel struct {
	DB *sql.DB
}

func (n NodeModel) Insert(node *node.Node) error {
	return nil
}

func (n NodeModel) Get(node *node.Node) error {
	return nil
}

func (n NodeModel) Update(node *node.Node) error {
	return nil
}

func (n NodeModel) Delete(id uint64) error {
	return nil
}
