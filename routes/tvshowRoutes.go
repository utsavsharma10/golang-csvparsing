package routes

import (
	"github.com/gorilla/mux"
)

func tvShowRouter(r *mux.Router, paramControl func()) {

	// Route handles & endpoints
	r.HandleFunc("/tvshows", controller.paramControl).Methods("GET")
}
