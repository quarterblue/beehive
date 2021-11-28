package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Health check for the API
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Job routes
	router.HandlerFunc(http.MethodPost, "/v1/jobs", app.createJobHandler)
	router.HandlerFunc(http.MethodGet, "/v1/jobs/:id", app.showJobHandler)
	router.HandlerFunc(http.MethodGet, "/v1/jobs/:id/start", app.startJobHandler)
	router.HandlerFunc(http.MethodGet, "/v1/jobs/:id/stop", app.startJobHandler)

	// Node routes
	router.HandlerFunc(http.MethodPost, "/v1/nodes", app.createNodeHandler)
	router.HandlerFunc(http.MethodPut, "/v1/nodes/:id", app.editNodeHandler)
	router.HandlerFunc(http.MethodGet, "/v1/nodes/:id", app.showNodeHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/nodes/:id", app.deleteNodeHandler)

	// Settings routes
	router.HandlerFunc(http.MethodPost, "/v1/setting", app.createSettingHandler)
	router.HandlerFunc(http.MethodPut, "/v1/setting", app.editSettingHandler)
	router.HandlerFunc(http.MethodGet, "/v1/setting", app.showSettingHandler)

	// Return the httprouter instance.
	return router
}
