package main

import (
	"app/pkg/config"
	"app/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	router := chi.NewRouter()

	router.Use(handleSession)
	router.Use(middleware.Logger)
	router.Use(CSRFcheck)
	router.Use(middleware.Recoverer)

	router.Get("/", http.HandlerFunc(handlers.Home))

	router.Get("/about", http.HandlerFunc(handlers.About))

	return router
}
