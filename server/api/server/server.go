package server

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/huydnt1801/chuyende/internal/config"
	"github.com/huydnt1801/chuyende/pkg/log"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	logger       logr.Logger
	accSrv       *AccountServer
	tripSrv      *TripServer
	sessionStore sessions.Store
}

type Option func(*Server)

func NewServer(accSrv *AccountServer, tripSrv *TripServer, opts ...Option) (*Server, error) {
	cfg := config.MustParseConfig()
	srv := &Server{
		logger:       log.ZapLogger(),
		accSrv:       accSrv,
		tripSrv:      tripSrv,
		sessionStore: CookieStore(cfg.SecretKey),
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
	e.Use(session.Middleware(s.sessionStore))
	e.Use(middleware.CORS())

	healthG := e.Group("/healthz")
	healthG.GET("/ready", ready)
	healthG.GET("/liveness", liveness)

	api := e.Group("/api/v1")
	accountG := api.Group("/accounts")
	accountG.POST("/phone", s.accSrv.CheckPhone)
	accountG.POST("/register", s.accSrv.Register)
	accountG.POST("/register/confirm", s.accSrv.RegisterConfirm)
	accountG.POST("/login", s.accSrv.Login)
	accountG.POST("/login/driver", s.accSrv.LoginDriver)

	tripG := api.Group("/trips")
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

func CookieStore(secret string) *sessions.CookieStore {
	sessionStore := sessions.NewCookieStore([]byte(secret))
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.SameSite = http.SameSiteNoneMode
	sessionStore.Options.Secure = true
	return sessionStore
}
