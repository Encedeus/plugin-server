package services

import (
	"context"
	"fmt"
	db2 "github.com/Encedeus/pluginServer/db"
	"github.com/Encedeus/pluginServer/ent"
	"github.com/Encedeus/pluginServer/ent/user"
	"github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/hashing"
	"github.com/Encedeus/pluginServer/proto"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/validate"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"strings"
	"time"
)

func CreateUser(ctx context.Context, db *ent.Client, req *protoapi.UserRegisterRequest) (*ent.User, error) {
	err := validate.IsUsername(req.Name)
	if err != nil {
		return nil, err
	}

	err = validate.IsEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = validate.IsPassword(req.Password)
	if err != nil {
		return nil, err
	}

	userData, err := db.User.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetPassword(hashing.HashPassword(req.Password)).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			if strings.Contains(err.Error(), user.FieldEmail) {
				return nil, errors.ErrEmailAlreadyTaken
			}
			if strings.Contains(err.Error(), user.FieldName) {
				return nil, errors.ErrUsernameAlreadyTaken
			}
		}

		log.Errorf("user creation failed: %v", err)
		return nil, errors.ErrQueryFailed
	}

	return userData, nil
}
func updateUserUsername(ctx context.Context, user ent.User, username string) error {
	err := validate.IsUsername(username)

	if err != nil {
		return err
	}

	fmt.Println(user.Name, username)

	if user.Name == username {
		return errors.ErrNewUsernameEqualsOld
	}

	_, err = user.Update().SetName(username).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.ErrUsernameAlreadyTaken
		}
		return errors.ErrQueryFailed
	}

	return nil
}

func updateUserEmail(ctx context.Context, user ent.User, email string) error {
	err := validate.IsEmail(email)
	if err != nil {
		return err
	}

	if user.Email == email {
		return errors.ErrNewEmailEqualsOld
	}

	_, err = user.Update().SetEmail(email).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.ErrEmailAlreadyTaken
		}
		return errors.ErrQueryFailed
	}

	return nil
}

func updateUserPassword(ctx context.Context, user ent.User, password string) error {

	err := validate.IsPassword(password)
	if err != nil {
		return err
	}

	if hashing.VerifyHash(password, user.Password) {
		return errors.ErrNewPasswordEqualsOld
	}

	_, err = user.Update().SetPassword(hashing.HashPassword(password)).Save(ctx)
	if err != nil {
		return errors.ErrQueryFailed
	}

	return nil
}

// UpdateUser updates the user given an updateInfo dto unhashed password needs to be provided
func UpdateUser(ctx context.Context, db *ent.Client, userId uuid.UUID, req *protoapi.UserUpdateRequest) (*ent.User, error) {
	if req.String() == "" {
		return nil, errors.ErrNotModified
	}

	userData, err := db.User.Query().Where(user.IDEQ(userId)).First(ctx)

	if err != nil {
		return nil, errors.ErrQueryFailed
	}

	if IsUserDeleted(userData) {
		return nil, errors.ErrUserDeleted
	}

	if req.Name != "" {
		err = updateUserUsername(ctx, *userData, req.Name)

		if err != nil {
			return nil, err
		}
	}
	if req.Email != "" {
		err = updateUserEmail(ctx, *userData, req.Email)

		if err != nil {
			return nil, err
		}
	}
	if req.Password != "" {
		err = updateUserPassword(ctx, *userData, req.Password)

		if err != nil {
			return nil, err
		}
	}

	updatedUser, err := db.User.Query().
		Where(user.IDEQ(userId)).
		Select(user.FieldName, user.FieldEmail, user.FieldUpdatedAt).
		First(ctx)

	return updatedUser, err
}

func DeleteUser(ctx context.Context, db *ent.Client, userId uuid.UUID) error {
	userData, err := db.User.Query().Where(user.IDEQ(userId)).First(ctx)
	if err != nil {
		return errors.ErrQueryFailed
	}

	if IsUserDeleted(userData) {
		return errors.ErrUserDeleted
	}

	userData, err = userData.Update().SetDeletedAt(time.Now()).Save(ctx)

	if err != nil {
		return errors.ErrQueryFailed
	}

	return nil
}

func FindOneUser(ctx context.Context, db *ent.Client, req *protoapi.UserFindOneRequest) (*protoapi.UserFindOneResponse, error) {
	userData, err := db.User.Query().
		Where(user.IDEQ(proto.ProtoUUIDToUUID(req.UserId))).
		Select(user.FieldID, user.FieldName, user.FieldCreatedAt, user.FieldUpdatedAt).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrUserNotFound
		}

		return nil, errors.ErrQueryFailed
	}

	if IsUserDeleted(userData) {
		return nil, errors.ErrUserDeleted
	}

	resp := &protoapi.UserFindOneResponse{
		User: proto.EntUserEntityToProtoUser(userData),
	}

	return resp, nil
}

// GetUser is not to be used to fulfil requests, rather use FindOneUser
func GetUser(ctx context.Context, db *ent.Client, userId uuid.UUID) (*ent.User, error) {
	userData, err := db.User.Query().
		Where(user.IDEQ(userId)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrUserNotFound
		}

		return nil, errors.ErrQueryFailed
	}

	if IsUserDeleted(userData) {
		return nil, errors.ErrUserDeleted
	}

	return userData, nil
}

func DoesUserWithUUIDExist(ctx context.Context, db *ent.Client, userId uuid.UUID) bool {
	userData, err := db.User.Query().Where(user.IDEQ(userId)).First(ctx)

	if err != nil || IsUserDeleted(userData) {
		return false
	}

	return true
}

func GetLastUpdate(ctx context.Context, db *ent.Client, userId uuid.UUID) (int64, error) {
	userData, err := db.User.Query().
		Where(user.IDEQ(userId), user.DeletedAtIsNil()).
		Select("updated_at").
		First(ctx)

	if err != nil {
		return 0, nil
	}

	return userData.UpdatedAt.Unix(), nil
}
func IsUserDeleted(userData *ent.User) bool {
	return userData.DeletedAt.Unix() != -62135596800
}

func CanUserBeAuthorized(ctx context.Context, tokenData TokenClaims) (bool, error) {
	db := db2.GetDb()

	userData, err := db.User.Get(ctx, tokenData.UserId)

	if err != nil {
		return false, errors.ErrQueryFailed
	}

	if userData.DeletedAt.Unix() != -62135596800 {
		return false, nil
	}

	if userData.AuthUpdatedAt.Unix() > tokenData.IssuedAt.Unix() {
		return false, nil
	}

	return true, nil
}

func IsUserUpdated(ctx context.Context, db *ent.Client, userId uuid.UUID, issuedAt int64) (bool, error) {
	lastUpdate, err := GetLastUpdate(ctx, db, userId)
	if err != nil {
		return true, err
	}

	if lastUpdate > issuedAt {
		return true, nil
	}

	return false, nil
}

func VerifyUserEmail(ctx context.Context, db *ent.Client, userId uuid.UUID) error {
	err := db.User.Update().SetEmailVerified(true).Where(user.ID(userId)).Exec(ctx)

	if err != nil {

		if ent.IsNotFound(err) {
			log.Error("impossible condition; something went very wrong; nonexistent user acquired auth")
		}

		return errors.ErrQueryFailed
	}

	return nil
}
