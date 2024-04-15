package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	// mux.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("root."))
	// })

	mux.Post("/", app.Broker)

	// http.ListenAndServe(":3333", r)
	return mux
}
