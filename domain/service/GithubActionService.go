package service

import (
	"go-cobraapi/adapter/http/request"
	"go-cobraapi/adapter/http/response"
	"go-cobraapi/domain"
	"go-cobraapi/domain/port/output"
	"os"
)

type GitActionService struct {
	repo     output.GitActionInterface
	username string
	token    string
}

func NewGitActionService(gr output.GitActionInterface) *GitActionService {
	return &GitActionService{
		repo:     gr,
		username: os.Getenv("GITHUB_USER"),
		token:    os.Getenv("GITHUB_API_KEY"),
	}
}

func (grs *GitActionService) Fetch(req request.GitActionFetch) ([]response.GitActionResponse, error) {
	var GitActionDomain []response.GitActionResponse
	repositories, err := grs.repo.Fetch(domain.GitRepo{Name: req.Name})
	if err != nil {
		return nil, err
	}

	for _, repo := range repositories {
		GitActionDomain = append(GitActionDomain, response.GitActionResponse(repo))
	}
	return GitActionDomain, nil
}
func (grs *GitActionService) Run(req request.GitActionRun) error {
	err := grs.repo.Run(domain.GitRepo{Name: req.Name, Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
