package main

import (
	"github.com/spf13/viper"
	"gitlab-variables/src/app"
	"gitlab-variables/src/cmd"
)

func init() {
	viper.BindEnv("GITLAB_ENV")
	app.ReadConfig(viper.GetString("GITLAB_ENV"))
}

func main() {
	cmdRepo := cmd.NewCommandRepo()
	//cmdRepo.AddUpdateCmd(comp)
	cmdRepo.Root.Execute()
}
