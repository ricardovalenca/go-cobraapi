package server

import (
	"fmt"

	"go-cobraapi/adapter/git/repository"
	"go-cobraapi/adapter/http/controller"
	"go-cobraapi/domain/service"

	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	router := gin.Default()

	GithubRepoRepository := repository.NewGithubRepoRepository()
	//GitLabRepository := repository.NewGitlabRepoRepository()
	GitRepoService := service.NewGitRepoService(GithubRepoRepository)
	GitRepoController := controller.NewGitRepoController(GitRepoService)

	GithubActionRepository := repository.NewGithubActionRepository()
	GitActionService := service.NewGitActionService(GithubActionRepository)
	GitActionController := controller.NewGitActionController(GitActionService)

	v1 := router.Group("/v1")
	{
		git := v1.Group("/git")
		{
			repo := git.Group("/repo")
			{
				repo.GET("/fetch", GitRepoController.Fetch)
			}
			action := git.Group("/action")
			{
				action.GET("/fetch", GitActionController.Fetch)
				action.GET("/run", GitActionController.Run)
			}
		}
	}

	if err := router.Run(":8080"); err != nil {
		panic(fmt.Sprintf("Error opening server: %v", err))
	}
}
