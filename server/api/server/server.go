package server

import (
	"github.com/go-logr/logr"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/huydnt1801/chuyende/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	logger logr.Logger
	accSrv *AccountServer
}

type Option func(*Server)

func NewServer(accSrv *AccountServer, opts ...Option) (*Server, error) {
	srv := &Server{
		logger: log.ZapLogger(),
		accSrv: accSrv,
	}
	for _, opt := range opts {
		opt(srv)
	}
	return srv, nil
}

func (s *Server) initRoutes() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = CustomHTTPErrorHandler(s.logger)
	e.Use(middleware.CORS())

	healthG := e.Group("/healthz")
	healthG.GET("/ready", ready)
	healthG.GET("/liveness", liveness)

	accountG := e.Group("/accounts")
	accountG.POST("/register", s.accSrv.register)
	accountG.POST("/register/confirm", s.accSrv.registerConfirm)
	return e
}

func (s *Server) Serve() error {
	e := s.initRoutes()
	return e.Start(":5000")
}
