package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab-variables/src/app"
	"gitlab-variables/src/list"
	"log"
)

func (c *CommandRepo) AddInitCmd(compound *list.Compound) {
	c.Root.AddCommand(initCmd(compound))
}

func initCmd(compound *list.Compound) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize project",
		Long:  `Initialize project. Add directories, create files, fetch variables from Gitlab and backup them.`,
		Args:  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configName := args[0]
			if compound.IsValidConfigName(configName) == false {
				log.Fatal("Invalid config file name")
				return
			}

			viper.Set("projectName", configName)
			app.ReadConfig(configName)

			compound.Init()
		},
	}
}
