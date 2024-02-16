package api

import (
	"encoding/json"
	"fmt"
	"github.com/Encedeus/pluginServer/errors"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

var headers = map[string]string{
	"User-Agent":    "EncedeusRegistry",
	"Authorization": "",
}

type asset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
}
type getByTagResponse struct {
	Assets []asset `json:"assets"`
}
type objectType struct {
	Sha  string `json:"sha"`
	Type string `json:"type"`
}
type getReleaseRefResponse struct {
	Object objectType `json:"object"`
}

func DoesGitHubRepoExist(repo *protoapi.GithubRepo) bool {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", repo.Username, repo.RepoName)
	req, _ := http.NewRequest("GET", url, nil)

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true
	}

	if res.StatusCode != 404 {
		log.Errorf("github api returned an error: %d", res.StatusCode)
	}

	return false
}

func DoesReleaseTagExistInRepo(repo *protoapi.GithubRepo, tag string) bool {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/tags/%s", repo.Username, repo.RepoName, tag)
	req, _ := http.NewRequest("GET", url, nil)

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true
	}

	if res.StatusCode != 404 {
		log.Errorf("github api returned an error: %d", res.StatusCode)
	}

	return false
}

func GetReleaseFileURI(repo *protoapi.GithubRepo, tag string) (*string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/tags/%s", repo.Username, repo.RepoName, tag)
	req, _ := http.NewRequest("GET", url, nil)

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body getByTagResponse
	json.NewDecoder(res.Body).Decode(&body)

	//fmt.Println(err, body)

	if res.StatusCode != 200 {
		return nil, errors.ErrGithubApiFailed
	}

	for _, asset := range body.Assets {
		if asset.Name == "Golf2.zip" { // todo: change to encPlugin.zip / .gz
			return &asset.DownloadURL, nil
		}
	}

	return nil, errors.ErrMissingPluginFile
}

func GetCommitSHAOfReleaseRef(repo *protoapi.GithubRepo, releaseTag string) (sha string, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/ref/tags/%s", repo.Username, repo.RepoName, releaseTag)
	req, _ := http.NewRequest("GET", url, nil)

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body getReleaseRefResponse
	json.NewDecoder(res.Body).Decode(&body)

	if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			return "", errors.ErrRepoGone
		}
		log.Errorf("unexpected status: %d", res.StatusCode)
		return "", errors.ErrGithubApiFailed
	}

	return body.Object.Sha, nil
}

func GetReadme(repo *protoapi.GithubRepo, releaseTag string) (string, error) {
	sha, err := GetCommitSHAOfReleaseRef(repo, releaseTag)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/README.md", repo.Username, repo.RepoName, sha)
	req, _ := http.NewRequest("GET", url, nil)

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.ErrNoReadme
	}

	bodyBytes, _ := io.ReadAll(res.Body)

	return string(bodyBytes), nil
}
