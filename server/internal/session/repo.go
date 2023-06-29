package session

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/huydnt1801/chuyende/internal/ent"
	"github.com/huydnt1801/chuyende/internal/ent/session"
	"github.com/huydnt1801/chuyende/pkg/log"
)

type SessionRepo interface {
	GetSession(ctx context.Context, sessionID string) (*ent.User, *ent.VehicleDriver, error)
	CreateSession(ctx context.Context, userID int, driverID int) (string, error)
	RevokeSession(ctx context.Context, sessionID string) error
}

type RepoImpl struct {
	logger logr.Logger
	client *ent.Client
}

func NewRepo(client *ent.Client) *RepoImpl {
	return &RepoImpl{
		logger: log.ZapLogger(),
		client: client,
	}
}

func (s *RepoImpl) GetSession(ctx context.Context, sessionID string) (*ent.Session, error) {
	sess, err := s.client.Session.Query().WithDriver().WithUser().
		Where(session.SessionID(sessionID)).Only(ctx)
	return sess, err
}

func (s *RepoImpl) CreateSession(ctx context.Context, userID int, driverID int) (*ent.Session, error) {
	sessionID := uuid.New().String()
	q := s.client.Session.Create().
		SetSessionID(sessionID).
		SetExpireIn(time.Now().Add(24 * 30 * time.Hour))

	if userID != 0 {
		q.SetUserID(userID)
	}
	if driverID != 0 {
		q.SetDriverID(driverID)
	}
	return q.Save(ctx)
}

func (s *RepoImpl) RevokeSession(ctx context.Context, sessionID string) error {
	_, err := s.client.Session.Update().
		Where(session.SessionID(sessionID)).
		SetRevoked(true).
		Save(ctx)
	return err
}
