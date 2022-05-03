package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab-variables/src/app"
	"gitlab-variables/src/list"
)

func (c *CommandRepo) AddUpdateCmd(compound *list.Compound) {
	c.Root.AddCommand(updateCmd(compound))
}

func updateCmd(compound *list.Compound) *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update Variables",
		Long:  `Update Variables`,
		Args:  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set("projectName", args[0])
			app.ReadConfig(args[0])

			compound.Update()
		},
	}
}
