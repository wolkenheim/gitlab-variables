package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"gitlab-variables/src/app"
	"gitlab-variables/src/backup"
	"gitlab-variables/src/cmd"
	"gitlab-variables/src/gitlab"
	"gitlab-variables/src/list"
	"net/http"
	"time"
)

func init() {
	viper.BindEnv("GITLAB_ENV")
	app.ReadConfig(viper.GetString("GITLAB_ENV"))
}

func main() {
	gitlabService := gitlab.NewGitlabService(gitlab.NewApiClient(&http.Client{
		Timeout: time.Second * 10,
	}))
	backupService := backup.NewBackup(afero.NewOsFs())
	comp := list.NewCompound(gitlabService, backupService)

	cmdRepo := cmd.NewCommandRepo()
	cmdRepo.AddUpdateCmd(comp)
	cmdRepo.Root.Execute()
}
