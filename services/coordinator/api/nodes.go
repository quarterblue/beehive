package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/quarterblue/beehive/internal/node"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

func (app *application) createNodeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		IpAddr   string `json:"ipaddr"`
		Port     string `json:"port"`
		Jobcount int    `json:"job_count"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		log.Println(err)
		return
	}

	node := &node.Node{
		Name:     input.Name,
		IpAddr:   input.IpAddr,
		Port:     input.Port,
		JobCount: uint32(input.Jobcount),
	}

	err = app.models.Nodes.Insert(node)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/nodes/%d", node.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"node": node}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showNodeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	node, err := app.models.Nodes.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"node": node}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) editNodeHandler(w http.ResponseWriter, r *http.Request) {}

func (app *application) deleteNodeHandler(w http.ResponseWriter, r *http.Request) {

}

type NodeModel struct {
	DB *sql.DB
}

func (n NodeModel) Insert(node *node.Node) error {

	query := `
		INSERT INTO nodes (node_name, ipaddr, port, job_count)
		VALUES ($1, $2, $3, $4)
		RETURNING node_id`

	args := []interface{}{node.Name, node.IpAddr, node.Port, node.JobCount}

	return n.DB.QueryRow(query, args...).Scan(&node.ID)
}

func (n NodeModel) Get(id int64) (*node.Node, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT node_id, node_name, ipaddr, port, job_count
        FROM nodes 
        WHERE node_id = $1`

	var node node.Node

	err := n.DB.QueryRow(query, id).Scan(
		&node.ID,
		&node.Name,
		&node.IpAddr,
		&node.Port,
		&node.JobCount,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &node, nil
}

func (n NodeModel) Update(node *node.Node) error {
	return nil
}

func (n NodeModel) Delete(id uint64) error {
	return nil
}
