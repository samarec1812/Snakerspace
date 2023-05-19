package http

import (
	"net/http"
	"time"

	"github.com/samarec1812/Snakerspace/internal/app"

	"github.com/gorilla/mux"
)

func NewHTTPServer(port string, a app.App) *http.Server {
	handler := mux.NewRouter()
	s := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	AppRouter(handler, a)
	return s
}
