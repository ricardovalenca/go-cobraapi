package input

import (
	"go-cobraapi/adapter/http/request"
	"go-cobraapi/adapter/http/response"
)

type GitReposInterface interface {
	Fetch() ([]response.GitRepoResponse, error)
}

type GitActionInterface interface {
	Fetch(request.GitActionFetch) ([]response.GitActionResponse, error)
	Run(request.GitActionRun) error
}
