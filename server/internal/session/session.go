package session

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/huydnt1801/chuyende/internal/driver"
	"github.com/huydnt1801/chuyende/internal/ent"
	"github.com/huydnt1801/chuyende/internal/user"
	entutil "github.com/huydnt1801/chuyende/internal/utils/ent"
	"github.com/huydnt1801/chuyende/pkg/log"
)

type SessionService interface {
	GetSession(ctx context.Context, sessionID string) (*user.User, *driver.Driver, error)
	CreateSession(ctx context.Context, userID int, driverID int) (string, error)
	RevokeSession(ctx context.Context, sessionID string) error
}

func NewSessionService(db *sql.DB) *SessionServiceImpl {
	client := entutil.NewClientFromDB(db)
	s := &SessionServiceImpl{
		logger:    log.ZapLogger(),
		entClient: client,
	}
	return s
}

type SessionServiceImpl struct {
	logger    logr.Logger
	entClient *ent.Client
}

func (s *SessionServiceImpl) GetSession(ctx context.Context, sessionID string) (*user.User, *driver.Driver, error) {
	repo := NewRepo(s.entClient)
	sess, err := repo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting session: %w", err)
	}
	if sess.ExpireIn.Before(time.Now()) {
		return nil, nil, InvalidSessionError{}
	}
	if usr := sess.Edges.User; usr != nil {
		u, err := user.DecodeUser(usr)
		if err != nil {
			return nil, nil, err
		}
		u.Password = ""
		return u, nil, nil
	}
	if dr := sess.Edges.Driver; dr != nil {
		d, err := driver.DecodeDriver(dr)
		if err != nil {
			return nil, nil, err
		}
		d.Password = ""
		return nil, d, nil
	}
	return nil, nil, NotFoundSessionError{}
}

func (s *SessionServiceImpl) CreateSession(ctx context.Context, userID int, driverID int) (string, error) {
	repo := NewRepo(s.entClient)
	sess, err := repo.CreateSession(ctx, userID, driverID)
	if err != nil {
		return "", fmt.Errorf("failed creating session: %w", err)
	}
	return sess.SessionID, nil
}

func (s *SessionServiceImpl) RevokeSession(ctx context.Context, sessionID string) error {
	repo := NewRepo(s.entClient)
	return repo.RevokeSession(ctx, sessionID)
}
