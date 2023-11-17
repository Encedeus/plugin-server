package proto

import (
	"github.com/Encedeus/pluginServer/ent"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

func ProtoUUIDToUUID(id *protoapi.UUID) uuid.UUID {
	return uuid.MustParse(id.Value)
}

func UUIDToProtoUUID(id uuid.UUID) *protoapi.UUID {
	return &protoapi.UUID{
		Value: id.String(),
	}
}

func EntUserEntityToProtoUser(user *ent.User) *protoapi.User {
	return &protoapi.User{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ProtoUserToEntUserEntity(user *protoapi.User) *ent.User {
	userId, _ := uuid.Parse(user.Id)
	return &ent.User{
		ID:        userId,
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
		DeletedAt: user.DeletedAt.AsTime(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func MarshalControllerProtoResponseToJSON(c *echo.Context, okStatus int, message proto.Message) (err error) {
	json, err := protojson.Marshal(message)
	if err != nil {
		return (*c).JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return (*c).JSONBlob(okStatus, json)
}
