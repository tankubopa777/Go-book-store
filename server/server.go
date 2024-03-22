package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tansan/config"
	"tansan/modules/middleware/middlewareHandler"
	"tansan/modules/middleware/middlewareRepository"
	"tansan/modules/middleware/middlewareUsecase"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type (
	server struct {
		app *echo.Echo
		db *mongo.Client
		cfg *config.Config
		middleware middlewareHandler.MiddlewareHandlerService
	}
)

func newMiddleware(cfg *config.Config) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func (s * server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Println("Start service:", s.cfg.App.Name)
	<-quit
	log.Println("Shutting down service:", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed{
		log.Fatal("Failed to start server:", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app: echo.New(),
		db: db,
		cfg: cfg,
		middleware: newMiddleware(cfg),
	}
	
	// Basic middleware
	// Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout: 30 * time.Second, 
	}))

	// CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	switch s.cfg.App.Name {
	case "auth":
	case "user":
	case "book":
	case "useBooks":
	case "payment":
	}

	// Use tweleve factor app design
	// https://12factor.net/config

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pctx, quit)

	// Listen
	s.httpListening()

}