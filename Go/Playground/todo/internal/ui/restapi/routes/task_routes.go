package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/viniciusps01/todo/internal/ui/restapi/handler"
	"github.com/viniciusps01/todo/internal/ui/restapi/middleware"
)

func taskRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/", handler.CreateTaskHandler)
		r.Get("/", handler.ReadAllTasksHandler)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.ReadTaskHandler)
			r.Patch("/", handler.UpdateTaskHandler)
			r.Delete("/", handler.DeleteTaskHandler)
		})
	})

	return r
}
