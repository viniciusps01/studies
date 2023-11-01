package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/viniciusps01/internal/config"
	"github.com/viniciusps01/internal/ui/restapi/handler"
)

func SetUpRoutes(config *config.AppConfig) {
	handler.SetUpHandlers(config)
}

func All() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/tasks", taskRoutes())

	return r
}
