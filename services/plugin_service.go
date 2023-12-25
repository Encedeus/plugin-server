package services

import (
	"context"
	"github.com/Encedeus/pluginServer/api"
	"github.com/Encedeus/pluginServer/ent"
	"github.com/Encedeus/pluginServer/ent/plugin"
	"github.com/Encedeus/pluginServer/ent/publication"
	errors2 "github.com/Encedeus/pluginServer/errors"
	"github.com/Encedeus/pluginServer/proto"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/Encedeus/pluginServer/validate"
	"github.com/google/uuid"
	"strings"
)

func simplifyGithubUri(repoURL string) string {
	repoURL = strings.ReplaceAll(repoURL, "www.", "")
	repoPath := strings.ReplaceAll(repoURL, "https://github.com/", "")
	return repoPath
}

func CreatePlugin(ctx context.Context, db *ent.Client, req *protoapi.PluginCreateRequest, ownerId uuid.UUID) (*ent.Plugin, error) {
	err := validate.IsPluginName(req.Name)

	if !validate.IsGitHubURL(req.RepoUri) {
		return nil, errors2.ErrInvalidGithubURL
	}

	githubRepo := proto.GithubUriToProtoGithubRepo(req.RepoUri)

	err = validate.IsGithubRepo(githubRepo)

	if err != nil {
		return nil, err
	}

	sourceData, err := db.Source.Create().
		SetRepository(simplifyGithubUri(req.RepoUri)).
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

func FindPluginByName(ctx context.Context, db *ent.Client, pluginName string) (*protoapi.Plugin, error) {
	pluginData, err := db.Plugin.Query().
		Where(plugin.Name(pluginName)).
		WithSource().
		WithOwner().
		WithPublications().
		First(ctx)
	if err != nil {

		if ent.IsNotFound(err) {
			return nil, errors2.ErrPluginNotFound
		}

		return nil, errors2.ErrQueryFailed
	}

	return proto.EntPluginEntityToProtoPlugin(pluginData), nil
}

func GetPluginWithAllEdges(ctx context.Context, db *ent.Client, pluginId uuid.UUID) (*ent.Plugin, error) {
	pluginData, err := db.Plugin.Query().
		Where(plugin.ID(pluginId)).
		WithSource().
		WithOwner().
		WithPublications().
		First(ctx)
	if err != nil {

		if ent.IsNotFound(err) {
			return nil, errors2.ErrPluginNotFound
		}

		return nil, errors2.ErrQueryFailed
	}

	return pluginData, nil
}

func GetPluginWithSource(ctx context.Context, db *ent.Client, ownerId uuid.UUID, pluginId uuid.UUID) (*ent.Plugin, error) {
	pluginData, err := db.Plugin.Query().
		Where(
			plugin.OwnerID(ownerId),
			plugin.ID(pluginId),
		).WithSource().
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors2.ErrPluginNotFound
		}

		return nil, err
	}

	return pluginData, nil
}

func PublishRelease(ctx context.Context, db *ent.Client, req *protoapi.PluginPublishReleaseRequest, pluginData *ent.Plugin) error {
	err := validate.IsReleaseName(req.Name)

	if err != nil {
		return err
	}

	githubRepo := proto.SimpleGithubUriToProtoGithubRepo(pluginData.Edges.Source.Repository)

	if api.DoesReleaseTagExistInRepo(githubRepo, req.GithubReleaseTag) {
		return errors2.ErrReleaseTagDoesNotExist
	}

	uri, err := api.GetReleaseFileURI(githubRepo, req.GithubReleaseTag)
	if err != nil {
		return err
	}

	nameTaken, err := db.Publication.Query().Where(publication.PluginID(pluginData.ID), publication.Name(req.Name)).Exist(ctx)

	if nameTaken {
		return errors2.ErrReleaseNameAlreadyInUse
	}

	_, err = db.Publication.Create().
		SetName(req.Name).
		SetPlugin(pluginData).
		SetURIToFile(*uri).
		Save(ctx)

	if err != nil {
		return errors2.ErrQueryFailed
	}

	return nil
}

func DeprecateRelease(ctx context.Context, db *ent.Client, pluginId uuid.UUID, releaseName string) error {
	err := db.Publication.Update().Where(publication.Name(releaseName), publication.PluginID(pluginId)).SetIsDeprecated(true).Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errors2.ErrReleaseNotFound
		}

		return errors2.ErrQueryFailed
	}

	return nil
}
