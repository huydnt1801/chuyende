package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/huydnt1801/chuyende/internal/ent"
	"github.com/huydnt1801/chuyende/internal/ent/session"
	entutil "github.com/huydnt1801/chuyende/internal/utils/ent"
	"github.com/huydnt1801/chuyende/pkg/log"
)

type SessionService interface {
	CreateSession(ctx context.Context, user *User, driverID int) (string, error)
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

func (s *SessionServiceImpl) CreateSession(ctx context.Context, user *User, driverID int) (string, error) {
	sessionID := uuid.New().String()
	q := s.entClient.Session.Create().
		SetSessionID(sessionID).
		SetExpireIn(time.Now().Add(24 * 30 * time.Hour))

	if user != nil {
		q.SetUserID(user.ID)
	}
	if driverID != 0 {
		q.SetDriverID(driverID)
	}
	sess, err := q.Save(ctx)
	if err != nil {
		return "", fmt.Errorf("failed creating session: %w", err)
	}
	return sess.SessionID, nil
}

func (s *SessionServiceImpl) RevokeSession(ctx context.Context, sessionID string) error {
	_, err := s.entClient.Session.Update().
		Where(session.SessionID(sessionID)).
		SetRevoked(true).
		Save(ctx)
	return err
}
