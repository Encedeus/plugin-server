package validate

import (
	"github.com/Encedeus/pluginServer/api"
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/errors"
	"github.com/microcosm-cc/bluemonday"
	"net"
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
		return errors.ErrInvalidUsername
	}

	return nil
}

func isGitHubURL(repoURL string) bool {
	u, err := url.Parse(repoURL)
	if err != nil {
		return false
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return false
		}
	}
	return host == "github.com"
}

func IsGithubRepo(repoURL string) error {
	repoURL = strings.ReplaceAll(repoURL, "www.", "")

	if !isGitHubURL(repoURL) {
		return errors.ErrInvalidGithubURL
	}

	repoPath := strings.ReplaceAll(repoURL, "https://github.com/", "")

	if !api.DoesGitHubRepoExist(repoPath) {
		return errors.ErrGithubRepoDoesNotExist
	}

	return nil
}
