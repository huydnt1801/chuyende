package user

import (
	"context"
	"database/sql"
	"regexp"
	"time"

	"github.com/go-logr/logr"
	"github.com/huydnt1801/chuyende/internal/ent"
	entutil "github.com/huydnt1801/chuyende/internal/utils/ent"
	"github.com/huydnt1801/chuyende/pkg/log"
)

type Service interface {
	CreateUser(ctx context.Context, user *User, password string) (*User, time.Time, error)
	Confirm(ctx context.Context, user *User, otp string) error
	FindUser(ctx context.Context, params *UserParams) (*User, error)
	SendConfirmationToken(ctx context.Context, user *User) (time.Time, error)
}

type ServiceImpl struct {
	logger    logr.Logger
	entClient *ent.Client
	// passwordHasher      PasswordHasher
	otpSender OTPService
}

func NewService(db *sql.DB) *ServiceImpl {
	client := entutil.NewClientFromDB(db)
	s := &ServiceImpl{
		logger:    log.ZapLogger(),
		entClient: client,
		otpSender: NewOTPService(db),
	}
	return s
}

func (s *ServiceImpl) CreateUser(ctx context.Context, user *User, password string) (*User, time.Time, error) {
	repo := NewRepo(s.entClient)
	if !isValidPhoneNumber(user.PhoneNumber) {
		return nil, time.Time{}, InvalidPhoneError{}
	}
	// validate pass and hash
	u, err := repo.CreateUser(ctx, user)
	if err != nil {
		return nil, time.Time{}, err
	}
	var nextTime time.Time
	nextTime, err = s.otpSender.SendOTP(ctx, user.PhoneNumber)
	if err != nil {
		return nil, time.Time{}, err
	}
	return u, nextTime, nil
}

func (s *ServiceImpl) Confirm(ctx context.Context, usr *User, otp string) error {
	repo := NewRepo(s.entClient)
	ok, err := s.otpSender.VerifyOTP(ctx, usr, otp)
	if err != nil {
		return err
	}
	if !ok {
		return InvalidOTPError{}
	}

	confirmed := true
	err = repo.UpdateUser(ctx, usr, &UserUpdate{
		Confirmed: &confirmed,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) FindUser(ctx context.Context, params *UserParams) (*User, error) {
	repo := NewRepo(s.entClient)
	return repo.FindUser(ctx, params)
}

func (s *ServiceImpl) SendConfirmationToken(ctx context.Context, user *User) (time.Time, error) {
	return s.otpSender.SendOTP(ctx, user.PhoneNumber)
}

func isValidPhoneNumber(phoneNumber string) bool {
	const PhoneNumberPattern = "^0[0-9]{9}$"
	if ok, _ := regexp.MatchString(PhoneNumberPattern, phoneNumber); !ok {
		return false
	}
	return true
}
