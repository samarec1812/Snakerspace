package main

import (
	"log"

	"github.com/samarec1812/Snakerspace/internal/adapters/noterepo"
	"github.com/samarec1812/Snakerspace/internal/app"
	"github.com/samarec1812/Snakerspace/internal/ports/http"
)

const httpPort = ":80"

func main() {
	httpServer := http.NewHTTPServer(httpPort, app.NewApp(noterepo.New()))

	log.Printf("start server on: %s\n", httpPort)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen on %s", httpPort)
	}
}
