package cmd

import (
	"github.com/spf13/cobra"
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
		Args:  nil,
		Run: func(cmd *cobra.Command, args []string) {
			compound.Init()
		},
	}
}
