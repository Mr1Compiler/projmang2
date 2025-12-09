package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
	"github.com/mustafaameen91/project-managment/backend/internal/server/routes"
)

type Server struct {
	port      string
	router    *gin.Engine
	container *container.Container
}

func NewServer(port string, c *container.Container) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Register middlewares
	router.Use(CORSMiddleware())
	router.Use(LoggingMiddleware())

	return &Server{
		port:      port,
		router:    router,
		container: c,
	}
}

func (s *Server) RegisterRoutes() {
	routes.RegisterRoutes(s.router, s.container)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Run() error {
	log.Printf("Server starting on port %s", s.port)
	return s.router.Run("0.0.0.0:" + s.port)
}
