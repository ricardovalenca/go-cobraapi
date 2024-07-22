package controller

import (
	"fmt"
	"go-cobraapi/adapter/http/request"
	"go-cobraapi/domain/port/input"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GitActionController struct {
	service input.GitActionInterface
}

func NewGitActionController(GitActionService input.GitActionInterface) *GitActionController {
	return &GitActionController{
		service: GitActionService,
	}
}
func (grc *GitActionController) Fetch(c *gin.Context) {
	var req request.GitActionFetch
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := grc.service.Fetch(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
func (grc *GitActionController) Run(c *gin.Context) {
	var req request.GitActionRun
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := grc.service.Run(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "OK")
}
