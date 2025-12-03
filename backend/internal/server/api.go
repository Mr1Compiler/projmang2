package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	router *gin.Engine
}

func NewServer(port string) *Server {
	router := gin.Default()

	// Register middlewares
	router.Use(CORSMiddleware())
	router.Use(LoggingMiddleware())

	return &Server{
		port:   port,
		router: router,
	}
}

func (s *Server) RegisterRoutes() {
	// Register your routes here
	// s.router.GET("/health", handlers.HealthCheck)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Run() error {
	log.Printf("Server starting on port %s", s.port)
	return s.router.Run(":" + s.port)
}
