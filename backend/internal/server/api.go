package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	router := gin.New()
	router.Use(gin.Recovery())

	// Super-Permissive Manual OPTIONS Handler (The "Big Hammer" fix)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// CORS Setup - Must be first to handle OPTIONS requests correctly
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "https://chic-luck-production.up.railway.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins in production for now (or strictly control via env)
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// Register middlewares (order matters - timeout should be early)
	router.Use(TimeoutMiddleware(cfg.QueryTimeout))
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
