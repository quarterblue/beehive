package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/jobs", app.createJobHandler)
	router.HandlerFunc(http.MethodGet, "/v1/jobs/:id", app.showJobHandler)
	router.HandlerFunc(http.MethodPost, "/v1/nodes", app.createNodeHandler)
	router.HandlerFunc(http.MethodPut, "/v1/nodes/:id", app.editNodeHandler)
	router.HandlerFunc(http.MethodGet, "/v1/nodes/:id", app.showNodeHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/nodes/:id", app.deleteNodeHandler)

	// Return the httprouter instance.
	return router
}
