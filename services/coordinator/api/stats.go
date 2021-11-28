package api

import (
	"net/http"
)

// import (
// 	"sync"

// 	"github.com/manucorporat/stats"
// )

// var (
// 	ips        = stats.New()
// 	messages   = stats.New()
// 	users      = stats.New()
// 	mutexStats sync.RWMutex
// 	savedStats map[string]uint64
// )

// func connectedUsers() uint64 {
// 	connected := users.Get("connected") - users.Get("disconnected")
// 	if connected < 0 {
// 		return 0
// 	}
// 	return uint64(connected)
// }

// // Stats returns savedStats data.
// func Stats() map[string]uint64 {
// 	mutexStats.RLock()
// 	defer mutexStats.RUnlock()

// 	return savedStats
// }

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
