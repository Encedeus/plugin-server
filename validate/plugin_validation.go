package validate

import (
	"github.com/Encedeus/pluginServer/api"
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/errors"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"net/url"
	"strings"
)

func IsPluginName(pluginName string) error {
	if len(pluginName) > config.Config.Validation.MaxPluginNameLen {
		return errors.ErrPluginNameTooLong
	}
	if len(pluginName) < config.Config.Validation.MinPluginNameLen {
		return errors.ErrPluginNameTooShort
	}

	p := bluemonday.StrictPolicy()
	if s := p.Sanitize(pluginName); s != pluginName {
		return errors.ErrInvalidPluginName
	}

	return nil
}

func IsReleaseName(username string) error {
	if len(username) > config.Config.Validation.MaxReleaseNameLen {
		return errors.ErrReleaseUsernameTooLong
	}
	if len(username) < config.Config.Validation.MinReleaseNameLen {
		return errors.ErrReleaseUsernameTooShort
	}

	p := bluemonday.StrictPolicy()
	if s := p.Sanitize(username); s != username {
		return errors.ErrInvalidUsername
	}

	return nil
}

func IsGitHubURL(repoURL string) bool {

	u, err1 := url.Parse(repoURL)
	u, err2 := url.Parse(repoURL)
	if err1 != nil && err2 != nil {
		return false
	}
	host := u.Host
	if strings.Contains(host, ":") {
		return false
	}
	return host == "github.com" || host == "www.github.com"
}

func IsGithubRepo(repo *protoapi.GithubRepo) error {
	if !api.DoesGitHubRepoExist(repo) {
		return errors.ErrGithubRepoDoesNotExist
	}

	return nil
}
