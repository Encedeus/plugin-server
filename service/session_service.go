package service

import (
	"PluginServer/ent/verifysession"
	"context"
	"github.com/google/uuid"
)

func StartVerifySession(userId uuid.UUID) (uuid.UUID, error) {
	sessionData, err := Db.VerifySession.Create().SetUserID(userId).Save(context.Background())

	if err != nil {
		return uuid.UUID{}, err
	}

	return sessionData.Sid, err
}

func CloseVerifySession(sessionId uuid.UUID) error {
	sessionData, err := Db.VerifySession.Query().Where(verifysession.Sid(sessionId)).First(context.Background())
	if err != nil {
		return err
	}

	err = Db.VerifySession.DeleteOne(sessionData).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func CloseVerifySessionByUserId(userId uuid.UUID) error {
	sessionData, err := Db.VerifySession.Query().Where(verifysession.UserID(userId)).First(context.Background())
	if err != nil {
		return err
	}

	err = Db.VerifySession.DeleteOne(sessionData).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
