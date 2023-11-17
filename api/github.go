package api

import (
	"fmt"
	"net/http"
)

func DoesGitHubRepoExist(repo string) bool {
	url := fmt.Sprintf("https://api.github.com/repos/%s", repo)
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true
	}

	return false
}
