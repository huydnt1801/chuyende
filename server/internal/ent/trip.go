// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/huydnt1801/chuyende/internal/ent/trip"
	"github.com/huydnt1801/chuyende/internal/ent/user"
	"github.com/huydnt1801/chuyende/internal/ent/vehicledriver"
)

// Trip is the model entity for the Trip schema.
type Trip struct {
	config `json:"-" mapstructure:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// DriverID holds the value of the "driver_id" field.
	DriverID int `json:"driver_id,omitempty" mapstructure:",omitempty"`
	// StartX holds the value of the "start_x" field.
	StartX float64 `json:"start_x,omitempty"`
	// StartY holds the value of the "start_y" field.
	StartY float64 `json:"start_y,omitempty"`
	// StartLocation holds the value of the "start_location" field.
	StartLocation string `json:"start_location,omitempty"`
	// EndX holds the value of the "end_x" field.
	EndX float64 `json:"end_x,omitempty"`
	// EndY holds the value of the "end_y" field.
	EndY float64 `json:"end_y,omitempty"`
	// EndLocation holds the value of the "end_location" field.
	EndLocation string `json:"end_location,omitempty"`
	// Distance holds the value of the "distance" field.
	Distance float64 `json:"distance,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Type holds the value of the "type" field.
	Type trip.Type `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status trip.Status `json:"status,omitempty"`
	// Rate holds the value of the "rate" field.
	Rate int `json:"rate,omitempty" mapstructure:",omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TripQuery when eager-loading is set.
	Edges        TripEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TripEdges holds the relations/edges for other nodes in the graph.
type TripEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Driver holds the value of the driver edge.
	Driver *VehicleDriver `json:"driver,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DriverOrErr returns the Driver value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) DriverOrErr() (*VehicleDriver, error) {
	if e.loadedTypes[1] {
		if e.Driver == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vehicledriver.Label}
		}
		return e.Driver, nil
	}
	return nil, &NotLoadedError{edge: "driver"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Trip) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case trip.FieldStartX, trip.FieldStartY, trip.FieldEndX, trip.FieldEndY, trip.FieldDistance, trip.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case trip.FieldID, trip.FieldUserID, trip.FieldDriverID, trip.FieldRate:
			values[i] = new(sql.NullInt64)
		case trip.FieldStartLocation, trip.FieldEndLocation, trip.FieldType, trip.FieldStatus:
			values[i] = new(sql.NullString)
		case trip.FieldCreatedAt, trip.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Trip fields.
func (t *Trip) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case trip.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case trip.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case trip.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case trip.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				t.UserID = int(value.Int64)
			}
		case trip.FieldDriverID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field driver_id", values[i])
			} else if value.Valid {
				t.DriverID = int(value.Int64)
			}
		case trip.FieldStartX:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field start_x", values[i])
			} else if value.Valid {
				t.StartX = value.Float64
			}
		case trip.FieldStartY:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field start_y", values[i])
			} else if value.Valid {
				t.StartY = value.Float64
			}
		case trip.FieldStartLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_location", values[i])
			} else if value.Valid {
				t.StartLocation = value.String
			}
		case trip.FieldEndX:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field end_x", values[i])
			} else if value.Valid {
				t.EndX = value.Float64
			}
		case trip.FieldEndY:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field end_y", values[i])
			} else if value.Valid {
				t.EndY = value.Float64
			}
		case trip.FieldEndLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field end_location", values[i])
			} else if value.Valid {
				t.EndLocation = value.String
			}
		case trip.FieldDistance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field distance", values[i])
			} else if value.Valid {
				t.Distance = value.Float64
			}
		case trip.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				t.Price = value.Float64
			}
		case trip.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = trip.Type(value.String)
			}
		case trip.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = trip.Status(value.String)
			}
		case trip.FieldRate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate", values[i])
			} else if value.Valid {
				t.Rate = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Trip.
// This includes values selected through modifiers, order, etc.
func (t *Trip) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Trip entity.
func (t *Trip) QueryUser() *UserQuery {
	return NewTripClient(t.config).QueryUser(t)
}

// QueryDriver queries the "driver" edge of the Trip entity.
func (t *Trip) QueryDriver() *VehicleDriverQuery {
	return NewTripClient(t.config).QueryDriver(t)
}

// Update returns a builder for updating this Trip.
// Note that you need to call Trip.Unwrap() before calling this method if this Trip
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Trip) Update() *TripUpdateOne {
	return NewTripClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Trip entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Trip) Unwrap() *Trip {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Trip is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Trip) String() string {
	var builder strings.Builder
	builder.WriteString("Trip(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteString(", ")
	builder.WriteString("driver_id=")
	builder.WriteString(fmt.Sprintf("%v", t.DriverID))
	builder.WriteString(", ")
	builder.WriteString("start_x=")
	builder.WriteString(fmt.Sprintf("%v", t.StartX))
	builder.WriteString(", ")
	builder.WriteString("start_y=")
	builder.WriteString(fmt.Sprintf("%v", t.StartY))
	builder.WriteString(", ")
	builder.WriteString("start_location=")
	builder.WriteString(t.StartLocation)
	builder.WriteString(", ")
	builder.WriteString("end_x=")
	builder.WriteString(fmt.Sprintf("%v", t.EndX))
	builder.WriteString(", ")
	builder.WriteString("end_y=")
	builder.WriteString(fmt.Sprintf("%v", t.EndY))
	builder.WriteString(", ")
	builder.WriteString("end_location=")
	builder.WriteString(t.EndLocation)
	builder.WriteString(", ")
	builder.WriteString("distance=")
	builder.WriteString(fmt.Sprintf("%v", t.Distance))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", t.Price))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("rate=")
	builder.WriteString(fmt.Sprintf("%v", t.Rate))
	builder.WriteByte(')')
	return builder.String()
}

// Trips is a parsable slice of Trip.
type Trips []*Trip
