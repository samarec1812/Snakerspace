package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/samarec1812/Snakerspace/config"
	"github.com/samarec1812/Snakerspace/internal/metrics"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samarec1812/Snakerspace/internal/adapters/noterepo"
	"github.com/samarec1812/Snakerspace/internal/app"
	_http "github.com/samarec1812/Snakerspace/internal/ports/http"
)

const (
// httpPort    = ":8080"
// metricsPort = ":8081"
)

var (
	configPath = flag.String("config", "", "File with config app")
)

func main() {
	flag.Parse()

	configs, err := config.InitConfig(*configPath)
	if err != nil {
		log.Fatalf("error with init config: %s", err.Error())
	}

	httpServer := _http.NewHTTPServer(configs.PortHTTP, app.NewApp(noterepo.New()))

	eg, ctx := errgroup.WithContext(context.Background())

	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			log.Printf("captured signal: %v\n", s)
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	eg.Go(func() error {
		log.Printf("starting http server, listening on %s\n", httpServer.Addr)
		defer log.Printf("close http server listening on %s\n", httpServer.Addr)

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s: %s", httpServer.Addr, err.Error())
			}

			close(errCh)
		}()

		go func() {
			if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		go func() {
			if err := metrics.Listen(configs.MetricsHTTP); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	})

	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}

	log.Println("servers were successfully shutdown")
}
