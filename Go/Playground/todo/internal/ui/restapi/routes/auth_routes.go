package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/viniciusps01/todo/internal/ui/restapi/handler"
	"github.com/viniciusps01/todo/internal/ui/restapi/middleware"
)

func authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUserHandler)

		r.Route("/", func(r chi.Router) {
			r.Use(middleware.Auth)
			r.Get("/", handler.ReadAllUsersHandler)
			r.Delete("/", handler.DeleteUserHandler)
		})

		r.Route("/profile", func(r chi.Router) {
			r.Use(middleware.Auth)
			r.Get("/", handler.ReadUserHandler)
		})
	})

	r.Post("/signin", handler.AuthenticateUserHandler)

	return r
}
