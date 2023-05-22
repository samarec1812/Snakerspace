package http

import (
	"github.com/gorilla/mux"
	"github.com/samarec1812/Snakerspace/internal/app"
	"github.com/sirupsen/logrus"
	"log"
)

func AppRouter(r *mux.Router, a app.App) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	lw := NewLoggerMiddleware(log.New(logger.Writer(), "", 0)) // log.New(os.Stdout, "", log.Ldate|log.Ltime))
	r.Use(lw.LogMiddleware)
	r.Use(lw.RecoverMiddleware)

	r.HandleFunc("/note", createNote(a)).Methods("POST")
	r.HandleFunc("/note", getAllNote(a)).Methods("GET")
	r.HandleFunc("/note/{id}", getNote(a)).Methods("GET")
	r.HandleFunc("/note/{id}", updateNote(a)).Methods("PUT")
	r.HandleFunc("/note/{id}", deleteNote(a)).Methods("DELETE")
}
