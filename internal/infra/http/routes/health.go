package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uzimbra/orch-process-transactions/internal/infra/http/handlers"
)

func HealthCheckRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/check", handlers.HealthCheckHandler)

	return r
}
