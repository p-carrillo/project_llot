package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/diteria/project_llot/backend/internal/adapters/config"
	httpadapter "github.com/diteria/project_llot/backend/internal/adapters/http"
	"github.com/diteria/project_llot/backend/internal/adapters/health"
	"github.com/diteria/project_llot/backend/internal/adapters/ingest/nginxjson"
	"github.com/diteria/project_llot/backend/internal/adapters/storage/jsonl"
	apphealth "github.com/diteria/project_llot/backend/internal/application/health"
	apptraffic "github.com/diteria/project_llot/backend/internal/application/traffic"
)

func main() {
	cfg := config.FromEnv()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))

	checker := health.NewStaticChecker(true)
	healthService := apphealth.NewService(checker)
	logParser := nginxjson.NewParser()
	trafficRepo, err := jsonl.NewRepository(cfg.DataFilePath)
	if err != nil {
		logger.Error("failed to initialize storage", "error", err, "data_file", cfg.DataFilePath)
		os.Exit(1)
	}
	trafficService := apptraffic.NewService(logParser, trafficRepo, cfg.SessionTimeout)
	handler := httpadapter.NewHandler(healthService, trafficService, cfg.MaxBodyBytes)
	server := httpadapter.NewServer(cfg, handler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		logger.Info("starting agent", "http_addr", cfg.HTTPAddr)
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		logger.Info("shutdown signal received")
	case serverErr := <-errCh:
		if serverErr != nil && serverErr != http.ErrServerClosed {
			logger.Error("server failed", "error", serverErr)
			os.Exit(1)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "error", err)
		os.Exit(1)
	}
	logger.Info("agent stopped")
}
