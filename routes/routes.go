package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"simpleserver.com/m/controllers"
)

// RegisterRoutes sets up the routes
func RegisterRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/api/notes", controllers.GetNoteHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/notes", controllers.PostNoteHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/notes/{id}", controllers.PutNoteHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNoteHandler).Methods(http.MethodDelete)

	return r
}
