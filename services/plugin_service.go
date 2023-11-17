package services

import (
	"context"
	"github.com/Encedeus/pluginServer/ent"
	errors2 "github.com/Encedeus/pluginServer/errors"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/validate"
	"github.com/google/uuid"
)

func CreatePlugin(ctx context.Context, db *ent.Client, req *protoapi.PluginCreateRequest, ownerId uuid.UUID) (*ent.Plugin, error) {
	err := validate.IsPluginName(req.Name)
	err = validate.IsGithubRepo(req.GithubRepo)

	if err != nil {
		return nil, err
	}

	sourceData, err := db.Source.Create().
		SetRepository(req.GithubRepo).
		Save(ctx)

	if err != nil {
		return nil, errors2.ErrQueryFailed
	}

	pluginData, err := db.Plugin.Create().
		SetName(req.Name).
		SetOwnerID(ownerId).
		SetSource(sourceData).
		Save(ctx)

	if err != nil {

		if ent.IsConstraintError(err) {
			db.Source.DeleteOne(sourceData).Exec(ctx)

			return nil, errors2.ErrPluginNameAlreadyTaken
		}

		return nil, errors2.ErrQueryFailed
	}

	return pluginData, nil
}
