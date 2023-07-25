package service

import (
	"PluginServer/dto"
	"PluginServer/ent"
	"PluginServer/ent/user"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

func CreateUser(name string, email string, passwordHash string) (*uuid.UUID, error) {
	userData, err := Db.User.Create().
		SetName(name).
		SetEmail(email).
		SetPassword(passwordHash).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return &userData.UUID, err
}

// UpdateUser updates the user given an updateInfo dto
func UpdateUser(updateInfo dto.UpdateUserDTO, userId uuid.UUID) error {
	userData, err := Db.User.Query().Where(user.UUID(userId)).First(context.Background())

	if err != nil {
		return err
	}

	if IsUserDeleted(userData) {
		return errors.New("user deleted")
	}

	if updateInfo.Name != "" {
		_, err = userData.Update().SetName(updateInfo.Name).Save(context.Background())
	}

	if updateInfo.Password != "" {
		_, err = userData.Update().SetPassword(updateInfo.Password).SetAuthUpdatedAt(time.Now()).Save(context.Background())
	}

	if updateInfo.Email != "" {
		_, err = userData.Update().SetEmail(updateInfo.Email).SetVerified(false).Save(context.Background())
	}

	return err
}

func DeleteUser(userId uuid.UUID) error {
	userData, err := Db.User.Query().Where(user.UUID(userId)).First(context.Background())
	if err != nil {
		return err
	}

	if IsUserDeleted(userData) {
		return errors.New("already deleted")
	}

	userData, err = userData.Update().SetDeletedAt(time.Now()).Save(context.Background())
	if err != nil {
		return err
	}

	return err
}

func GetUser(userId uuid.UUID) (*ent.User, error) {
	userData, err := Db.User.Query().
		Where(user.UUID(userId)).
		Select("uuid", "name", "created_at", "updated_at", "deleted_at", "email").
		First(context.Background())
	if err != nil {
		return nil, err
	}

	if IsUserDeleted(userData) {
		return nil, errors.New("user deleted")
	}

	return userData, err
}

func DoesUserWithUUIDExist(userId uuid.UUID) bool {
	userData, err := Db.User.Query().Where(user.UUID(userId)).First(context.Background())

	if err != nil || IsUserDeleted(userData) {
		return false
	}

	return true
}

func GetLastAuthUpdate(userId uuid.UUID) (int64, error) {
	userData, err := Db.User.Query().
		Where(user.UUID(userId), user.DeletedAtIsNil()).
		Select("auth_updated_at").
		First(context.Background())

	if err != nil {
		return 0, err
	}

	return userData.AuthUpdatedAt.Unix(), nil
}
func IsUserDeleted(userData *ent.User) bool {
	return userData.DeletedAt.Unix() != -62135596800
}

func VerifyUserEmail(userId uuid.UUID) error {
	_, err := Db.User.Update().Where(user.UUID(userId)).SetVerified(true).Save(context.Background())
	return err
}
