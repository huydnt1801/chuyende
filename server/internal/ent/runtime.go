// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/huydnt1801/chuyende/internal/ent/otp"
	"github.com/huydnt1801/chuyende/internal/ent/schema"
	"github.com/huydnt1801/chuyende/internal/ent/session"
	"github.com/huydnt1801/chuyende/internal/ent/trip"
	"github.com/huydnt1801/chuyende/internal/ent/user"
	"github.com/huydnt1801/chuyende/internal/ent/vehicledriver"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	otpFields := schema.Otp{}.Fields()
	_ = otpFields
	// otpDescCreatedAt is the schema descriptor for created_at field.
	otpDescCreatedAt := otpFields[2].Descriptor()
	// otp.DefaultCreatedAt holds the default value on creation for the created_at field.
	otp.DefaultCreatedAt = otpDescCreatedAt.Default.(func() time.Time)
	sessionMixin := schema.Session{}.Mixin()
	sessionMixinFields0 := sessionMixin[0].Fields()
	_ = sessionMixinFields0
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionMixinFields0[0].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionMixinFields0[1].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sessionDescRevoked is the schema descriptor for revoked field.
	sessionDescRevoked := sessionFields[4].Descriptor()
	// session.DefaultRevoked holds the default value on creation for the revoked field.
	session.DefaultRevoked = sessionDescRevoked.Default.(bool)
	tripMixin := schema.Trip{}.Mixin()
	tripMixinFields0 := tripMixin[0].Fields()
	_ = tripMixinFields0
	tripFields := schema.Trip{}.Fields()
	_ = tripFields
	// tripDescCreatedAt is the schema descriptor for created_at field.
	tripDescCreatedAt := tripMixinFields0[0].Descriptor()
	// trip.DefaultCreatedAt holds the default value on creation for the created_at field.
	trip.DefaultCreatedAt = tripDescCreatedAt.Default.(func() time.Time)
	// tripDescUpdatedAt is the schema descriptor for updated_at field.
	tripDescUpdatedAt := tripMixinFields0[1].Descriptor()
	// trip.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	trip.DefaultUpdatedAt = tripDescUpdatedAt.Default.(func() time.Time)
	// trip.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	trip.UpdateDefaultUpdatedAt = tripDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tripDescRate is the schema descriptor for rate field.
	tripDescRate := tripFields[8].Descriptor()
	// trip.RateValidator is a validator for the "rate" field. It is called by the builders before save.
	trip.RateValidator = func() func(int) error {
		validators := tripDescRate.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(rate int) error {
			for _, fn := range fns {
				if err := fn(rate); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescPhoneNumber is the schema descriptor for phone_number field.
	userDescPhoneNumber := userFields[0].Descriptor()
	// user.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	user.PhoneNumberValidator = userDescPhoneNumber.Validators[0].(func(string) error)
	// userDescConfirmed is the schema descriptor for confirmed field.
	userDescConfirmed := userFields[1].Descriptor()
	// user.DefaultConfirmed holds the default value on creation for the confirmed field.
	user.DefaultConfirmed = userDescConfirmed.Default.(bool)
	vehicledriverMixin := schema.VehicleDriver{}.Mixin()
	vehicledriverMixinFields0 := vehicledriverMixin[0].Fields()
	_ = vehicledriverMixinFields0
	vehicledriverFields := schema.VehicleDriver{}.Fields()
	_ = vehicledriverFields
	// vehicledriverDescCreatedAt is the schema descriptor for created_at field.
	vehicledriverDescCreatedAt := vehicledriverMixinFields0[0].Descriptor()
	// vehicledriver.DefaultCreatedAt holds the default value on creation for the created_at field.
	vehicledriver.DefaultCreatedAt = vehicledriverDescCreatedAt.Default.(func() time.Time)
	// vehicledriverDescUpdatedAt is the schema descriptor for updated_at field.
	vehicledriverDescUpdatedAt := vehicledriverMixinFields0[1].Descriptor()
	// vehicledriver.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vehicledriver.DefaultUpdatedAt = vehicledriverDescUpdatedAt.Default.(func() time.Time)
	// vehicledriver.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vehicledriver.UpdateDefaultUpdatedAt = vehicledriverDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vehicledriverDescPhoneNumber is the schema descriptor for phone_number field.
	vehicledriverDescPhoneNumber := vehicledriverFields[0].Descriptor()
	// vehicledriver.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	vehicledriver.PhoneNumberValidator = vehicledriverDescPhoneNumber.Validators[0].(func(string) error)
}
