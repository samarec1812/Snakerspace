package http

import (
	"github.com/samarec1812/Snakerspace/internal/app"

	"github.com/gorilla/mux"
)

func AppRouter(r *mux.Router, a app.App) {

	r.HandleFunc("/note", createNote(a)).Methods("POST")
	r.HandleFunc("/note", getAllNote(a)).Methods("GET")
	r.HandleFunc("/note/{id}", getNote(a)).Methods("GET")
	r.HandleFunc("/note/{id}", updateNote(a)).Methods("PUT")
	r.HandleFunc("/note/{id}", deleteNote(a)).Methods("DELETE")
}
