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
	logger  logr.Logger
	accSrv  *AccountServer
	tripSrv *TripServer
}

type Option func(*Server)

func NewServer(accSrv *AccountServer, tripSrv *TripServer, opts ...Option) (*Server, error) {
	srv := &Server{
		logger:  log.ZapLogger(),
		accSrv:  accSrv,
		tripSrv: tripSrv,
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
	accountG.POST("/phone", s.accSrv.CheckPhone)
	accountG.POST("/register", s.accSrv.Register)
	accountG.POST("/register/confirm", s.accSrv.RegisterConfirm)
	accountG.POST("/login", s.accSrv.Login)

	tripG := e.Group("/trips")
	tripG.GET("", s.tripSrv.ListTrip)
	tripG.POST("", s.tripSrv.CreateTrip)
	tripG.PATCH("/:tripId/status", s.tripSrv.UpdateStatusTrip)
	tripG.PATCH("/:tripId/rate", s.tripSrv.RateTrip)
	tripG.PATCH("/:tripId/accept", s.tripSrv.AcceptTrip)
	return e
}

func (s *Server) Serve() error {
	e := s.initRoutes()
	return e.Start(":5000")
}
