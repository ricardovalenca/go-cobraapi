package repository

import (
	"encoding/json"
	"fmt"
	"go-cobraapi/domain"
	"net/http"
	"os"
)

type GithubRepoRepository struct {
}

func NewGithubRepoRepository() *GithubRepoRepository {
	return &GithubRepoRepository{}
}

func (grs *GithubRepoRepository) Fetch() ([]domain.GitRepo, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", os.Getenv("GITHUB_USER"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: Status code %d", resp.StatusCode)
	}

	var repositories []domain.GitRepo
	decoder := json.NewDecoder(resp.Body)
	//body, err := io.ReadAll(resp.Body)
	//fmt.Println("Body:", string(body))
	if err := decoder.Decode(&repositories); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return repositories, nil
}
