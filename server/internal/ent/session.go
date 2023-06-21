// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/huydnt1801/chuyende/internal/ent/session"
	"github.com/huydnt1801/chuyende/internal/ent/user"
	"github.com/huydnt1801/chuyende/internal/ent/vehicledriver"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// DriverID holds the value of the "driver_id" field.
	DriverID int `json:"driver_id,omitempty"`
	// SessionID holds the value of the "session_id" field.
	SessionID string `json:"session_id,omitempty"`
	// ExpireIn holds the value of the "expire_in" field.
	ExpireIn time.Time `json:"expire_in,omitempty"`
	// Revoked holds the value of the "revoked" field.
	Revoked bool `json:"revoked,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges        SessionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
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
func (e SessionEdges) UserOrErr() (*User, error) {
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
func (e SessionEdges) DriverOrErr() (*VehicleDriver, error) {
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
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldRevoked:
			values[i] = new(sql.NullBool)
		case session.FieldID, session.FieldUserID, session.FieldDriverID:
			values[i] = new(sql.NullInt64)
		case session.FieldSessionID:
			values[i] = new(sql.NullString)
		case session.FieldCreatedAt, session.FieldUpdatedAt, session.FieldExpireIn:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case session.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case session.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = int(value.Int64)
			}
		case session.FieldDriverID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field driver_id", values[i])
			} else if value.Valid {
				s.DriverID = int(value.Int64)
			}
		case session.FieldSessionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field session_id", values[i])
			} else if value.Valid {
				s.SessionID = value.String
			}
		case session.FieldExpireIn:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire_in", values[i])
			} else if value.Valid {
				s.ExpireIn = value.Time
			}
		case session.FieldRevoked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field revoked", values[i])
			} else if value.Valid {
				s.Revoked = value.Bool
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Session.
// This includes values selected through modifiers, order, etc.
func (s *Session) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Session entity.
func (s *Session) QueryUser() *UserQuery {
	return NewSessionClient(s.config).QueryUser(s)
}

// QueryDriver queries the "driver" edge of the Session entity.
func (s *Session) QueryDriver() *VehicleDriverQuery {
	return NewSessionClient(s.config).QueryDriver(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return NewSessionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("driver_id=")
	builder.WriteString(fmt.Sprintf("%v", s.DriverID))
	builder.WriteString(", ")
	builder.WriteString("session_id=")
	builder.WriteString(s.SessionID)
	builder.WriteString(", ")
	builder.WriteString("expire_in=")
	builder.WriteString(s.ExpireIn.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("revoked=")
	builder.WriteString(fmt.Sprintf("%v", s.Revoked))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session
