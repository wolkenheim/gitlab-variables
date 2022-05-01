package main

import (
	"gitlab-variables/src/app"
	"gitlab-variables/src/cmd"
)

func init() {
	app.ReadConfig("gitlab_one")
}

func main() {
	cmdRepo := cmd.NewCommandRepo()
	//cmdRepo.AddUpdateCmd(comp)
	cmdRepo.Root.Execute()
}
