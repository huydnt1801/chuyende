// Code generated by ent, DO NOT EDIT.

package vehicledriver

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/huydnt1801/chuyende/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldUpdatedAt, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldPhoneNumber, v))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldFullName, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldPassword, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLTE(FieldUpdatedAt, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldContainsFold(FieldFullName, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldContainsFold(FieldPassword, v))
}

// LicenseEQ applies the EQ predicate on the "license" field.
func LicenseEQ(v License) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldEQ(FieldLicense, v))
}

// LicenseNEQ applies the NEQ predicate on the "license" field.
func LicenseNEQ(v License) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNEQ(FieldLicense, v))
}

// LicenseIn applies the In predicate on the "license" field.
func LicenseIn(vs ...License) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldIn(FieldLicense, vs...))
}

// LicenseNotIn applies the NotIn predicate on the "license" field.
func LicenseNotIn(vs ...License) predicate.VehicleDriver {
	return predicate.VehicleDriver(sql.FieldNotIn(FieldLicense, vs...))
}

// HasTrips applies the HasEdge predicate on the "trips" edge.
func HasTrips() predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TripsTable, TripsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTripsWith applies the HasEdge predicate on the "trips" edge with a given conditions (other predicates).
func HasTripsWith(preds ...predicate.Trip) predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		step := newTripsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSessions applies the HasEdge predicate on the "sessions" edge.
func HasSessions() predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionsWith applies the HasEdge predicate on the "sessions" edge with a given conditions (other predicates).
func HasSessionsWith(preds ...predicate.Session) predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		step := newSessionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VehicleDriver) predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VehicleDriver) predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VehicleDriver) predicate.VehicleDriver {
	return predicate.VehicleDriver(func(s *sql.Selector) {
		p(s.Not())
	})
}
