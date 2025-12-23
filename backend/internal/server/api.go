package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/config"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
	"github.com/mustafa91ameen/prjalgo/backend/internal/server/routes"
)

type Server struct {
	port      string
	router    *gin.Engine
	container *container.Container
	httpSrv   *http.Server
}

func NewServer(port string, c *container.Container, cfg *config.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Register middlewares (order matters - timeout should be early)
	router.Use(TimeoutMiddleware(cfg.QueryTimeout))
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
	s.httpSrv = &http.Server{
		Addr:    "0.0.0.0:" + s.port,
		Handler: s.router,
	}
	log.Printf("Server starting on port %s", s.port)
	return s.httpSrv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpSrv == nil {
		return nil
	}
	return s.httpSrv.Shutdown(ctx)
}
