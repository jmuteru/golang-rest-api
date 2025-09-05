package main

import (
	"log"
	"net/http"
	"time"
)

type Application struct {
	config Config
}
type Config struct {
	addr string
}

// Router setup
func (app *Application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	return mux
}

// Start server logic
func (app *Application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addr, // application port
		Handler:      mux,             //Define default handler
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server running on port %s", app.config.addr)
	return srv.ListenAndServe()
}
