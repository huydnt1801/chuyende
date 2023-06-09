// Code generated by ent, DO NOT EDIT.

package trip

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/huydnt1801/chuyende/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldUserID, v))
}

// DriverID applies equality check predicate on the "driver_id" field. It's identical to DriverIDEQ.
func DriverID(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldDriverID, v))
}

// StartX applies equality check predicate on the "start_x" field. It's identical to StartXEQ.
func StartX(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStartX, v))
}

// StartY applies equality check predicate on the "start_y" field. It's identical to StartYEQ.
func StartY(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStartY, v))
}

// StartLocation applies equality check predicate on the "start_location" field. It's identical to StartLocationEQ.
func StartLocation(v string) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStartLocation, v))
}

// EndX applies equality check predicate on the "end_x" field. It's identical to EndXEQ.
func EndX(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEndX, v))
}

// EndY applies equality check predicate on the "end_y" field. It's identical to EndYEQ.
func EndY(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEndY, v))
}

// EndLocation applies equality check predicate on the "end_location" field. It's identical to EndLocationEQ.
func EndLocation(v string) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEndLocation, v))
}

// Distance applies equality check predicate on the "distance" field. It's identical to DistanceEQ.
func Distance(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldDistance, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldPrice, v))
}

// Rate applies equality check predicate on the "rate" field. It's identical to RateEQ.
func Rate(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldRate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldUserID, vs...))
}

// DriverIDEQ applies the EQ predicate on the "driver_id" field.
func DriverIDEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldDriverID, v))
}

// DriverIDNEQ applies the NEQ predicate on the "driver_id" field.
func DriverIDNEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldDriverID, v))
}

// DriverIDIn applies the In predicate on the "driver_id" field.
func DriverIDIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldDriverID, vs...))
}

// DriverIDNotIn applies the NotIn predicate on the "driver_id" field.
func DriverIDNotIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldDriverID, vs...))
}

// DriverIDIsNil applies the IsNil predicate on the "driver_id" field.
func DriverIDIsNil() predicate.Trip {
	return predicate.Trip(sql.FieldIsNull(FieldDriverID))
}

// DriverIDNotNil applies the NotNil predicate on the "driver_id" field.
func DriverIDNotNil() predicate.Trip {
	return predicate.Trip(sql.FieldNotNull(FieldDriverID))
}

// StartXEQ applies the EQ predicate on the "start_x" field.
func StartXEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStartX, v))
}

// StartXNEQ applies the NEQ predicate on the "start_x" field.
func StartXNEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldStartX, v))
}

// StartXIn applies the In predicate on the "start_x" field.
func StartXIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldStartX, vs...))
}

// StartXNotIn applies the NotIn predicate on the "start_x" field.
func StartXNotIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldStartX, vs...))
}

// StartXGT applies the GT predicate on the "start_x" field.
func StartXGT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldStartX, v))
}

// StartXGTE applies the GTE predicate on the "start_x" field.
func StartXGTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldStartX, v))
}

// StartXLT applies the LT predicate on the "start_x" field.
func StartXLT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldStartX, v))
}

// StartXLTE applies the LTE predicate on the "start_x" field.
func StartXLTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldStartX, v))
}

// StartYEQ applies the EQ predicate on the "start_y" field.
func StartYEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStartY, v))
}

// StartYNEQ applies the NEQ predicate on the "start_y" field.
func StartYNEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldStartY, v))
}

// StartYIn applies the In predicate on the "start_y" field.
func StartYIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldStartY, vs...))
}

// StartYNotIn applies the NotIn predicate on the "start_y" field.
func StartYNotIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldStartY, vs...))
}

// StartYGT applies the GT predicate on the "start_y" field.
func StartYGT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldStartY, v))
}

// StartYGTE applies the GTE predicate on the "start_y" field.
func StartYGTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldStartY, v))
}

// StartYLT applies the LT predicate on the "start_y" field.
func StartYLT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldStartY, v))
}

// StartYLTE applies the LTE predicate on the "start_y" field.
func StartYLTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldStartY, v))
}

// StartLocationEQ applies the EQ predicate on the "start_location" field.
func StartLocationEQ(v string) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStartLocation, v))
}

// StartLocationNEQ applies the NEQ predicate on the "start_location" field.
func StartLocationNEQ(v string) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldStartLocation, v))
}

// StartLocationIn applies the In predicate on the "start_location" field.
func StartLocationIn(vs ...string) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldStartLocation, vs...))
}

// StartLocationNotIn applies the NotIn predicate on the "start_location" field.
func StartLocationNotIn(vs ...string) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldStartLocation, vs...))
}

// StartLocationGT applies the GT predicate on the "start_location" field.
func StartLocationGT(v string) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldStartLocation, v))
}

// StartLocationGTE applies the GTE predicate on the "start_location" field.
func StartLocationGTE(v string) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldStartLocation, v))
}

// StartLocationLT applies the LT predicate on the "start_location" field.
func StartLocationLT(v string) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldStartLocation, v))
}

// StartLocationLTE applies the LTE predicate on the "start_location" field.
func StartLocationLTE(v string) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldStartLocation, v))
}

// StartLocationContains applies the Contains predicate on the "start_location" field.
func StartLocationContains(v string) predicate.Trip {
	return predicate.Trip(sql.FieldContains(FieldStartLocation, v))
}

// StartLocationHasPrefix applies the HasPrefix predicate on the "start_location" field.
func StartLocationHasPrefix(v string) predicate.Trip {
	return predicate.Trip(sql.FieldHasPrefix(FieldStartLocation, v))
}

// StartLocationHasSuffix applies the HasSuffix predicate on the "start_location" field.
func StartLocationHasSuffix(v string) predicate.Trip {
	return predicate.Trip(sql.FieldHasSuffix(FieldStartLocation, v))
}

// StartLocationEqualFold applies the EqualFold predicate on the "start_location" field.
func StartLocationEqualFold(v string) predicate.Trip {
	return predicate.Trip(sql.FieldEqualFold(FieldStartLocation, v))
}

// StartLocationContainsFold applies the ContainsFold predicate on the "start_location" field.
func StartLocationContainsFold(v string) predicate.Trip {
	return predicate.Trip(sql.FieldContainsFold(FieldStartLocation, v))
}

// EndXEQ applies the EQ predicate on the "end_x" field.
func EndXEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEndX, v))
}

// EndXNEQ applies the NEQ predicate on the "end_x" field.
func EndXNEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldEndX, v))
}

// EndXIn applies the In predicate on the "end_x" field.
func EndXIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldEndX, vs...))
}

// EndXNotIn applies the NotIn predicate on the "end_x" field.
func EndXNotIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldEndX, vs...))
}

// EndXGT applies the GT predicate on the "end_x" field.
func EndXGT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldEndX, v))
}

// EndXGTE applies the GTE predicate on the "end_x" field.
func EndXGTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldEndX, v))
}

// EndXLT applies the LT predicate on the "end_x" field.
func EndXLT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldEndX, v))
}

// EndXLTE applies the LTE predicate on the "end_x" field.
func EndXLTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldEndX, v))
}

// EndYEQ applies the EQ predicate on the "end_y" field.
func EndYEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEndY, v))
}

// EndYNEQ applies the NEQ predicate on the "end_y" field.
func EndYNEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldEndY, v))
}

// EndYIn applies the In predicate on the "end_y" field.
func EndYIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldEndY, vs...))
}

// EndYNotIn applies the NotIn predicate on the "end_y" field.
func EndYNotIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldEndY, vs...))
}

// EndYGT applies the GT predicate on the "end_y" field.
func EndYGT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldEndY, v))
}

// EndYGTE applies the GTE predicate on the "end_y" field.
func EndYGTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldEndY, v))
}

// EndYLT applies the LT predicate on the "end_y" field.
func EndYLT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldEndY, v))
}

// EndYLTE applies the LTE predicate on the "end_y" field.
func EndYLTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldEndY, v))
}

// EndLocationEQ applies the EQ predicate on the "end_location" field.
func EndLocationEQ(v string) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldEndLocation, v))
}

// EndLocationNEQ applies the NEQ predicate on the "end_location" field.
func EndLocationNEQ(v string) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldEndLocation, v))
}

// EndLocationIn applies the In predicate on the "end_location" field.
func EndLocationIn(vs ...string) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldEndLocation, vs...))
}

// EndLocationNotIn applies the NotIn predicate on the "end_location" field.
func EndLocationNotIn(vs ...string) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldEndLocation, vs...))
}

// EndLocationGT applies the GT predicate on the "end_location" field.
func EndLocationGT(v string) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldEndLocation, v))
}

// EndLocationGTE applies the GTE predicate on the "end_location" field.
func EndLocationGTE(v string) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldEndLocation, v))
}

// EndLocationLT applies the LT predicate on the "end_location" field.
func EndLocationLT(v string) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldEndLocation, v))
}

// EndLocationLTE applies the LTE predicate on the "end_location" field.
func EndLocationLTE(v string) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldEndLocation, v))
}

// EndLocationContains applies the Contains predicate on the "end_location" field.
func EndLocationContains(v string) predicate.Trip {
	return predicate.Trip(sql.FieldContains(FieldEndLocation, v))
}

// EndLocationHasPrefix applies the HasPrefix predicate on the "end_location" field.
func EndLocationHasPrefix(v string) predicate.Trip {
	return predicate.Trip(sql.FieldHasPrefix(FieldEndLocation, v))
}

// EndLocationHasSuffix applies the HasSuffix predicate on the "end_location" field.
func EndLocationHasSuffix(v string) predicate.Trip {
	return predicate.Trip(sql.FieldHasSuffix(FieldEndLocation, v))
}

// EndLocationEqualFold applies the EqualFold predicate on the "end_location" field.
func EndLocationEqualFold(v string) predicate.Trip {
	return predicate.Trip(sql.FieldEqualFold(FieldEndLocation, v))
}

// EndLocationContainsFold applies the ContainsFold predicate on the "end_location" field.
func EndLocationContainsFold(v string) predicate.Trip {
	return predicate.Trip(sql.FieldContainsFold(FieldEndLocation, v))
}

// DistanceEQ applies the EQ predicate on the "distance" field.
func DistanceEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldDistance, v))
}

// DistanceNEQ applies the NEQ predicate on the "distance" field.
func DistanceNEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldDistance, v))
}

// DistanceIn applies the In predicate on the "distance" field.
func DistanceIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldDistance, vs...))
}

// DistanceNotIn applies the NotIn predicate on the "distance" field.
func DistanceNotIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldDistance, vs...))
}

// DistanceGT applies the GT predicate on the "distance" field.
func DistanceGT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldDistance, v))
}

// DistanceGTE applies the GTE predicate on the "distance" field.
func DistanceGTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldDistance, v))
}

// DistanceLT applies the LT predicate on the "distance" field.
func DistanceLT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldDistance, v))
}

// DistanceLTE applies the LTE predicate on the "distance" field.
func DistanceLTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldDistance, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldPrice, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldType, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldStatus, vs...))
}

// RateEQ applies the EQ predicate on the "rate" field.
func RateEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldEQ(FieldRate, v))
}

// RateNEQ applies the NEQ predicate on the "rate" field.
func RateNEQ(v int) predicate.Trip {
	return predicate.Trip(sql.FieldNEQ(FieldRate, v))
}

// RateIn applies the In predicate on the "rate" field.
func RateIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldIn(FieldRate, vs...))
}

// RateNotIn applies the NotIn predicate on the "rate" field.
func RateNotIn(vs ...int) predicate.Trip {
	return predicate.Trip(sql.FieldNotIn(FieldRate, vs...))
}

// RateGT applies the GT predicate on the "rate" field.
func RateGT(v int) predicate.Trip {
	return predicate.Trip(sql.FieldGT(FieldRate, v))
}

// RateGTE applies the GTE predicate on the "rate" field.
func RateGTE(v int) predicate.Trip {
	return predicate.Trip(sql.FieldGTE(FieldRate, v))
}

// RateLT applies the LT predicate on the "rate" field.
func RateLT(v int) predicate.Trip {
	return predicate.Trip(sql.FieldLT(FieldRate, v))
}

// RateLTE applies the LTE predicate on the "rate" field.
func RateLTE(v int) predicate.Trip {
	return predicate.Trip(sql.FieldLTE(FieldRate, v))
}

// RateIsNil applies the IsNil predicate on the "rate" field.
func RateIsNil() predicate.Trip {
	return predicate.Trip(sql.FieldIsNull(FieldRate))
}

// RateNotNil applies the NotNil predicate on the "rate" field.
func RateNotNil() predicate.Trip {
	return predicate.Trip(sql.FieldNotNull(FieldRate))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDriver applies the HasEdge predicate on the "driver" edge.
func HasDriver() predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DriverTable, DriverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDriverWith applies the HasEdge predicate on the "driver" edge with a given conditions (other predicates).
func HasDriverWith(preds ...predicate.VehicleDriver) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		step := newDriverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Trip) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Trip) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
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
func Not(p predicate.Trip) predicate.Trip {
	return predicate.Trip(func(s *sql.Selector) {
		p(s.Not())
	})
}
