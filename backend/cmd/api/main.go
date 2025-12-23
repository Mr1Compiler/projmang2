package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mustafa91ameen/prjalgo/backend/internal/config"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
	"github.com/mustafa91ameen/prjalgo/backend/internal/db"
	"github.com/mustafa91ameen/prjalgo/backend/internal/logger"
	"github.com/mustafa91ameen/prjalgo/backend/internal/server"
)

func main() {
	// Initialize logger
	logger.Init()

	// Load config from .env
	cfg := config.Load()

	// Connect to database with retry logic
	ctx := context.Background()
	pool, err := db.Connect(ctx, cfg.DSN(), cfg.DBMaxRetries, cfg.DBRetryDelay)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	log.Println("Connected to database")

	// Build dependencies
	c := container.New(pool, cfg)

	// Create and start server
	srv := server.NewServer(cfg.ServerPort, c, cfg)
	srv.RegisterRoutes()

	// Start server in goroutine
	errChan := make(chan error, 1)
	go func() {
		log.Printf("Server starting on port %s", cfg.ServerPort)
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("Server failed: %v", err)
	case sig := <-quit:
		log.Printf("Received signal: %v, shutting down gracefully...", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error: %v", err)
		}
		log.Println("Server stopped")
	}
}
