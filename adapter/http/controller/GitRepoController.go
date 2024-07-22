package controller

import (
	"go-cobraapi/domain/port/input"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GitRepoController struct {
	service input.GitReposInterface
}

func NewGitRepoController(GitRepoService input.GitReposInterface) *GitRepoController {
	return &GitRepoController{
		service: GitRepoService,
	}
}
func (grc *GitRepoController) Fetch(c *gin.Context) {
	response, err := grc.service.Fetch()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, response)
}
