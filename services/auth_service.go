package services

import (
	"context"
	"github.com/Encedeus/pluginServer/ent"
	"github.com/Encedeus/pluginServer/ent/user"
	"github.com/Encedeus/pluginServer/errors"
	"github.com/google/uuid"
)

// GetUserUUIDAndHashByUsername returns the user's UserId and hashed password provided the username of the user
func GetUserUUIDAndHashByUsername(ctx context.Context, db *ent.Client, username string) (string, *uuid.UUID, error) {
	userData, err := db.User.Query().Where(user.Name(username)).Select(user.FieldID, user.FieldPassword).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil, errors.ErrUserNotFound
		}

		return "", nil, errors.ErrQueryFailed
	}

	if IsUserDeleted(userData) {
		return "", nil, errors.ErrUserDeleted
	}

	return userData.Password, &userData.ID, nil
}

// GetUserAuthDataAndHashByEmail returns the user's UserId and hashed password provided the email of the user
func GetUserAuthDataAndHashByEmail(ctx context.Context, db *ent.Client, email string) (string, *uuid.UUID, error) {
	userData, err := db.User.Query().Where(user.Email(email)).Select(user.FieldID, user.FieldPassword).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil, errors.ErrUserNotFound
		}

		return "", nil, errors.ErrQueryFailed
	}

	if IsUserDeleted(userData) {
		return "", nil, errors.ErrUserDeleted
	}

	return userData.Password, &userData.ID, nil
}
