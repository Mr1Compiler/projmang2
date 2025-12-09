package main

import (
	"context"
	"log"

	"github.com/mustafaameen91/project-managment/backend/internal/config"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
	"github.com/mustafaameen91/project-managment/backend/internal/db"
	"github.com/mustafaameen91/project-managment/backend/internal/logger"
	"github.com/mustafaameen91/project-managment/backend/internal/server"
)

func main() {
	// Initialize logger
	logger.Init()

	// Load config from .env
	cfg := config.Load()

	// Connect to database
	ctx := context.Background()
	pool, err := db.Connect(ctx, cfg.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	log.Println("Connected to database")

	// Build dependencies
	c := container.New(pool, cfg)

	// Create and start server
	srv := server.NewServer(cfg.ServerPort, c)
	srv.RegisterRoutes()

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
