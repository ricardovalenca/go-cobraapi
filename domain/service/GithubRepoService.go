package service

import (
	"go-cobraapi/adapter/http/response"
	"go-cobraapi/domain/port/output"
	"os"
)

type GitRepoService struct {
	repo     output.GitReposInterface
	username string
	token    string
}

func NewGitRepoService(gr output.GitReposInterface) *GitRepoService {
	return &GitRepoService{
		repo:     gr,
		username: os.Getenv("GITHUB_USER"),
		token:    os.Getenv("GITHUB_API_KEY"),
	}
}

func (grs *GitRepoService) Fetch() ([]response.GitRepoResponse, error) {
	var GitRepoDomain []response.GitRepoResponse
	repositories, err := grs.repo.Fetch()
	if err != nil {
		return nil, err
	}
	for _, repo := range repositories {
		GitRepoDomain = append(GitRepoDomain, response.GitRepoResponse(repo))
	}
	return GitRepoDomain, nil
}
