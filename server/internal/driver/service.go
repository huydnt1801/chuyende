package driver

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"github.com/go-logr/logr"
	"github.com/huydnt1801/chuyende/internal/ent"
	entutil "github.com/huydnt1801/chuyende/internal/utils/ent"
	"github.com/huydnt1801/chuyende/pkg/log"
)

type Service interface {
	Authenticate(ctx context.Context, username, password string) (*Driver, error)
}

type ServiceImpl struct {
	logger    logr.Logger
	entClient *ent.Client
}

func NewService(db *sql.DB) *ServiceImpl {
	client := entutil.NewClientFromDB(db)
	s := &ServiceImpl{
		logger:    log.ZapLogger(),
		entClient: client,
	}
	return s
}

func (s *ServiceImpl) Authenticate(ctx context.Context, phoneNumber, password string) (*Driver, error) {
	repo := NewRepo(s.entClient)

	if !isValidPhoneNumber(phoneNumber) {
		return nil, InvalidPhoneError{}
	}
	user, err := repo.FindDriver(ctx, phoneNumber)
	if err != nil {
		if IsDriverNotFound(err) {
			return nil, DriverNotFoundError{}
		} else {
			return nil, fmt.Errorf("failed querying user from DB: %w", err)
		}

	}

	if password != user.Password {
		return nil, InvalidPasswordError{}
	}

	return user, nil
}

func isValidPhoneNumber(phoneNumber string) bool {
	const PhoneNumberPattern = "^0[0-9]{9}$"
	if ok, _ := regexp.MatchString(PhoneNumberPattern, phoneNumber); !ok {
		return false
	}
	return true
}