package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

// Router setup using Chi
func (app *Application) mount() *chi.Mux {
	r := chi.NewRouter()
	//Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)    //Logging middleware
	r.Use(middleware.Recoverer) //Recover from panics

	//Middleware timeout
	r.Use(middleware.Timeout(60 * time.Second))

	//Good practice: append api version to each endpoint
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	return r
}

// Start server logic
func (app *Application) run(mux *chi.Mux) error {
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
