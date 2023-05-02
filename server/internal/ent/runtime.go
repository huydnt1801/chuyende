// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/huydnt1801/chuyende/internal/ent/otp"
	"github.com/huydnt1801/chuyende/internal/ent/schema"
	"github.com/huydnt1801/chuyende/internal/ent/user"
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
}