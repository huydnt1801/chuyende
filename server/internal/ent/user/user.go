// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldConfirmed holds the string denoting the confirmed field in the database.
	FieldConfirmed = "confirmed"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeTrips holds the string denoting the trips edge name in mutations.
	EdgeTrips = "trips"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TripsTable is the table that holds the trips relation/edge.
	TripsTable = "trips"
	// TripsInverseTable is the table name for the Trip entity.
	// It exists in this package in order to avoid circular dependency with the "trip" package.
	TripsInverseTable = "trips"
	// TripsColumn is the table column denoting the trips relation/edge.
	TripsColumn = "user_id"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPhoneNumber,
	FieldConfirmed,
	FieldFullName,
	FieldImageURL,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// DefaultConfirmed holds the default value on creation for the "confirmed" field.
	DefaultConfirmed bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByConfirmed orders the results by the confirmed field.
func ByConfirmed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfirmed, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByTripsCount orders the results by trips count.
func ByTripsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTripsStep(), opts...)
	}
}

// ByTrips orders the results by trips terms.
func ByTrips(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTripsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySessionsCount orders the results by sessions count.
func BySessionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSessionsStep(), opts...)
	}
}

// BySessions orders the results by sessions terms.
func BySessions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSessionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTripsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TripsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TripsTable, TripsColumn),
	)
}
func newSessionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
	)
}
