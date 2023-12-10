package api

import (
	"encoding/json"
	"fmt"
	"github.com/Encedeus/pluginServer/errors"
	protoapi "github.com/Encedeus/pluginServer/proto/go"
	"github.com/labstack/gommon/log"
	"net/http"
)

type asset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
}
type getByTagResponse struct {
	Assets []asset `json:"assets"`
}

func DoesGitHubRepoExist(repo *protoapi.GithubRepo) bool {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", repo.Username, repo.RepoName)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "")

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

	req.Header.Add("Authorization", "")

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
	req.Header.Add("Authorization", "")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var body getByTagResponse
	json.NewDecoder(res.Body).Decode(&body)

	//fmt.Println(err, body)

	if res.StatusCode != 200 {
		return nil, errors.GithubAPIError
	}

	for _, asset := range body.Assets {
		if asset.Name == "Golf2.zip" {
			return &asset.DownloadURL, nil
		}
	}

	return nil, errors.ErrMissingPluginFile
}
