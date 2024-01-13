package proto

import (
	"github.com/Encedeus/pluginServer/ent"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strings"
)

func EntUserEntityToProtoUser(user *ent.User) *protoapi.User {
	return &protoapi.User{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func EntPublicationEntityToProtoRelease(publication *ent.Publication) *protoapi.Release {
	return &protoapi.Release{
		Name:         publication.Name,
		PublishedAt:  timestamppb.New(publication.CreatedAt),
		IsDeprecated: publication.IsDeprecated,
		DownloadURI:  publication.URIToFile,
	}
}

func EntPublicationsEntityToProtoReleases(publications []*ent.Publication) []*protoapi.Release {
	var releases []*protoapi.Release

	for _, publication := range publications {
		releases = append(releases, EntPublicationEntityToProtoRelease(publication))
	}

	return releases
}

func EntPluginEntityToProtoPlugin(plugin *ent.Plugin) *protoapi.Plugin {
	return &protoapi.Plugin{
		Id:        plugin.ID.String(),
		Name:      plugin.Name,
		OwnerName: plugin.Edges.Owner.Name,
		Source: &protoapi.Source{
			RepoUri: "https://github.com/" + plugin.Edges.Source.Repository,
		},
		Releases: EntPublicationsEntityToProtoReleases(plugin.Edges.Publications),
	}
}

func GithubUriToProtoGithubRepo(repoURL string) *protoapi.GithubRepo {

	repoURL = strings.ReplaceAll(repoURL, "www.", "")
	repoPath := strings.ReplaceAll(repoURL, "https://github.com/", "")
	repoPathSegments := strings.Split(repoPath, "/")

	return &protoapi.GithubRepo{Username: repoPathSegments[0], RepoName: repoPathSegments[1]}
}

func SimpleGithubUriToProtoGithubRepo(repoURL string) *protoapi.GithubRepo {
	repoPathSegments := strings.Split(repoURL, "/")

	return &protoapi.GithubRepo{Username: repoPathSegments[0], RepoName: repoPathSegments[1]}
}

func MarshalControllerProtoResponseToJSON(c *echo.Context, status int, message proto.Message) (err error) {
	json, err := protojson.Marshal(message)
	if err != nil {
		return (*c).JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return (*c).JSONBlob(status, json)
}
