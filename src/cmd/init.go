package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab-variables/src/app"
	"gitlab-variables/src/list"
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
			viper.Set("projectName", args[0])
			app.ReadConfig(args[0])

			compound.Init()
		},
	}
}
