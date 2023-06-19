package trip

import (
	"github.com/huydnt1801/chuyende/internal/ent/trip"
)

type Trip struct {
	ID      int         `json:"id"`
	UserID  int         `json:"userId"`
	DriveID string      `json:"driveId,omitempty" mapstructure:",omitempty"`
	StartX  float64     `json:"startX"`
	StartY  float64     `json:"startY"`
	EndX    float64     `json:"endX"`
	EndY    float64     `json:"endY"`
	Price   float64     `json:"price"`
	Status  trip.Status `json:"status,omitempty" mapstructure:",omitempty"`
	Rate    int         `json:"rate,omitempty" mapstructure:",omitempty"`
}

type TripParams struct {
	TripID  *int
	UserID  *int
	DriveID *int
	Status  *trip.Status
	Rate    *int
}

type TripUpdate struct {
	DriveID *int
	Status  *trip.Status
	Rate    *int
}
