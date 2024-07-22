package repository

import (
	"encoding/json"
	"fmt"
	"go-cobraapi/domain"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type GithubActionRepository struct {
}

func NewGithubActionRepository() *GithubActionRepository {
	return &GithubActionRepository{}
}

type WorkflowsResponse struct {
	TotalCount int                `json:"total_count"`
	Workflows  []domain.GitAction `json:"workflows"`
}

func (grs *GithubActionRepository) Fetch(repo domain.GitRepo) ([]domain.GitAction, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows", os.Getenv("GITHUB_USER"), repo.Name)
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

	decoder := json.NewDecoder(resp.Body)
	//body, err := io.ReadAll(resp.Body)
	//fmt.Println("Body:", string(body))
	var workflowsResponse WorkflowsResponse

	if err := decoder.Decode(&workflowsResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	fmt.Println(workflowsResponse.Workflows)
	return workflowsResponse.Workflows, nil
}

func (grs *GithubActionRepository) Run(repo domain.GitRepo) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%d/dispatches", os.Getenv("GITHUB_USER"), repo.Name, repo.Id)

	client := resty.New()

	// Payload for the workflow dispatch event
	payload := map[string]interface{}{
		"ref": "main", // The branch or tag name the workflow should run on
		"inputs": map[string]interface{}{
			// Add any input fields your workflow requires
			"example_input": "value",
		},
	}

	// Personal Access Token (PAT)
	token := os.Getenv("GITHUB_API_KEY")
	if token == "" {
		return fmt.Errorf("please set your GITHUB_TOKEN environment variable")
	}

	// Send POST request to trigger the workflow
	resp, err := client.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		return fmt.Errorf("error triggering workflow: %v", err)
	}

	// Check if the request was successful
	if resp.IsError() {
		return fmt.Errorf("failed to trigger workflow: %s", resp.String())
	}
	return nil
}
