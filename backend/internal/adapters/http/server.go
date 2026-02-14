package http

import (
	"net/http"
	"time"

	"github.com/diteria/project_llot/backend/internal/adapters/config"
)

func NewServer(cfg config.Config, handler Handler) *http.Server {
	return &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           handler.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
