package http

import (
	"encoding/json"
	"github.com/samarec1812/Snakerspace/internal/metrics"
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

type LoggerMiddleware struct {
	logger *log.Logger
}

func NewLoggerMiddleware(logger *log.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger: logger,
	}
}

func (l *LoggerMiddleware) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		raw := r.URL.RawQuery

		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)

		timeStamp := time.Now()
		latency := timeStamp.Sub(start)

		status := lrw.Status()

		defer metrics.ObserveRequest(time.Since(start), status)

		method := r.Method

		bodySize := lrw.Size()
		if raw != "" {
			path = path + "?" + raw
		}

		if status >= 500 {
			l.logger.SetPrefix("ERROR:\t")
		} else {
			l.logger.SetPrefix("INFO:\t")
		}

		l.logger.Printf("method: %s, path: %s, status_code: %d, body_size: %dB, latency: %s\n",
			method,
			path,
			status,
			bodySize,
			latency,
		)
	})
}

func (l *LoggerMiddleware) RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				l.logger.Println(err) // May be log this error? Send to sentry?

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
