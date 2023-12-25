package services

import (
	"context"
	"github.com/Encedeus/pluginServer/ent"
	"github.com/Encedeus/pluginServer/ent/verificationsession"
	"github.com/Encedeus/pluginServer/errors"
	"github.com/google/uuid"
)

func StartVerificationSession(ctx context.Context, db *ent.Client, userId uuid.UUID) (*ent.VerificationSession, error) {
	sessionData, err := db.VerificationSession.Create().SetUserID(userId).Save(ctx)
	if err != nil {
		return nil, errors.ErrQueryFailed
	}
	return sessionData, nil
}

func CloseVerificationSession(ctx context.Context, db *ent.Client, sessionId string) error {
	_, err := db.VerificationSession.Delete().Where(verificationsession.ID(sessionId)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrSessionNotFound
		}

		return errors.ErrQueryFailed
	}

	return nil
}
func CloseVerificationSessionByUserId(ctx context.Context, db *ent.Client, userId uuid.UUID) error {
	_, err := db.VerificationSession.Delete().Where(verificationsession.UserID(userId)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrSessionNotFound
		}

		return errors.ErrQueryFailed
	}

	return nil
}

func GetVerificationSessionById(ctx context.Context, db *ent.Client, sessionId string) (*ent.VerificationSession, error) {
	sessionData, err := db.VerificationSession.Query().Where(verificationsession.ID(sessionId)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrSessionNotFound
		}

		return nil, errors.ErrQueryFailed
	}

	return sessionData, nil
}
