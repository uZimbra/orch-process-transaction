package services

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/uzimbra/orch-process-transactions/internal/config"
	"github.com/uzimbra/orch-process-transactions/internal/infra/http/routes"
	"go.uber.org/zap"
)

func StartApi() {
	port := config.GetServerEnv().Port
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Mount("/health", routes.HealthCheckRouter())

	zap.L().Sugar().Infow("Server started, listening on", "port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		zap.L().Sugar().Fatalw("Server start failed", "error", err)
	}
}
