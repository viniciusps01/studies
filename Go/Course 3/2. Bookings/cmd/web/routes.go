package main

import (
	"app/internals/config"
	"app/internals/handlers"
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

	fs := http.FileServer(http.Dir("../../static"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.Get("/", http.HandlerFunc(handlers.Home))

	router.Get("/about", http.HandlerFunc(handlers.About))

	router.Get("/contact", http.HandlerFunc(handlers.Contact))

	router.Get("/search-availability", http.HandlerFunc(handlers.SearchAvailability))
	router.Post("/search-availability", http.HandlerFunc(handlers.SearchAvailabilityPost))
	router.Post("/search-availability-json", http.HandlerFunc(handlers.AvailabilityJson))

	router.Get("/make-reservation", http.HandlerFunc(handlers.MakeReservation))
	router.Post("/make-reservation", http.HandlerFunc(handlers.MakeReservationPost))
	router.Get("/reservation-summary", http.HandlerFunc(handlers.ReservationSummary))

	router.Get("/majors-suite", http.HandlerFunc(handlers.MajorsSuite))

	router.Get("/generals-quarters", http.HandlerFunc(handlers.GeneralsQuarters))

	return router
}
