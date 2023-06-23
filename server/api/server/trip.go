package server

import (
	"database/sql"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/huydnt1801/chuyende/internal/trip"
	"github.com/huydnt1801/chuyende/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type TripServer struct {
	logger logr.Logger
	svc    trip.Service
}

func NewTripServer(db *sql.DB) *TripServer {
	svc := trip.NewService(db)
	srv := &TripServer{
		logger: log.ZapLogger(),
		svc:    svc,
	}
	return srv
}

type ListTripRequest struct {
	TripID  *int    `query:"tripId"`
	UserID  *int    `query:"userId"`
	DriveID *int    `query:"driveId"`
	Status  *string `query:"status"`
	Rate    *int    `query:"rate"`
}

type ListTripResponse struct {
	Code int          `json:"code"`
	Data []*trip.Trip `json:"data"`
}

func (s *TripServer) ListTrip(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &ListTripRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	params := &trip.TripParams{}
	err := mapstructure.Decode(data, params)
	if err != nil {
		return err
	}
	trip, err := s.svc.FindTrip(ctx, params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ListTripResponse{Code: http.StatusOK, Data: trip})
}

type CreateTripRequest struct {
	UserID int     `json:"userId" validate:"required"`
	StartX float64 `json:"startX" validate:"required"`
	StartY float64 `json:"startY" validate:"required"`
	EndX   float64 `json:"endX" validate:"required"`
	EndY   float64 `json:"endY" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}

type CreateTripResponse struct {
	Code int        `json:"code"`
	Data *trip.Trip `json:"data"`
}

func (s *TripServer) CreateTrip(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &CreateTripRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	createParams := &trip.Trip{}
	err := mapstructure.Decode(data, createParams)
	if err != nil {
		return err
	}
	trip, err := s.svc.CreateTrip(ctx, createParams)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, CreateTripResponse{Code: http.StatusOK, Data: trip})
}

type UpdateStatusTripRequest struct {
	TripID int    `param:"tripId" validate:"required,numeric"`
	Status string `json:"status"`
}

type UpdateStatusTripResponse struct {
	Code int        `json:"code"`
	Data *trip.Trip `json:"data"`
}

func (s *TripServer) UpdateStatusTrip(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &UpdateStatusTripRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	updateParams := &trip.TripUpdate{}
	err := mapstructure.Decode(data, updateParams)
	if err != nil {
		return err
	}
	trip, err := s.svc.UpdateTrip(ctx, data.TripID, updateParams)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, UpdateStatusTripResponse{Code: http.StatusOK, Data: trip})
}

type RateTripRequest struct {
	TripID int `param:"tripId" validate:"required,numeric"`
	Rate   int `json:"rate"`
}

type RateTripResponse struct {
	Code int        `json:"code"`
	Data *trip.Trip `json:"data"`
}

func (s *TripServer) RateTrip(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &RateTripRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	updateParams := &trip.TripUpdate{}
	err := mapstructure.Decode(data, updateParams)
	if err != nil {
		return err
	}
	trip, err := s.svc.UpdateTrip(ctx, data.TripID, updateParams)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, RateTripResponse{Code: http.StatusOK, Data: trip})
}

type AcceptTripRequest struct {
	TripID int `param:"tripId" validate:"required,numeric"`
}

type AcceptTripResponse struct {
	Code int        `json:"code"`
	Data *trip.Trip `json:"data"`
}

func (s *TripServer) AcceptTrip(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &AcceptTripRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	updateParams := &trip.TripUpdate{}
	err := mapstructure.Decode(data, updateParams)
	if err != nil {
		return err
	}
	trip, err := s.svc.UpdateTrip(ctx, data.TripID, updateParams)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, AcceptTripResponse{Code: http.StatusOK, Data: trip})
}
