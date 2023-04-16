package server

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Server struct {
	// svr string
}

type Option func(*Server)

func NewServer(opts ...Option) (*Server, error) {
	srv := &Server{}
	for _, opt := range opts {
		opt(srv)
	}
	return srv, nil
}

func (s *Server) initRoutes() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Binder = &CustomBinder{}

	healthG := e.Group("/healthz")
	healthG.GET("/ready", ready)
	healthG.GET("/liveness", liveness)

	return e
}

func (s *Server) Serve() error {
	e := s.initRoutes()
	return e.Start(":5000")
}
