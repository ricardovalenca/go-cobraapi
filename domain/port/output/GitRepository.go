package output

import (
	"go-cobraapi/domain"
)

type GitReposInterface interface {
	Fetch() ([]domain.GitRepo, error)
}

type GitActionInterface interface {
	Fetch(domain.GitRepo) ([]domain.GitAction, error)
	Run(domain.GitRepo) error
}
