package cmd

import (
	"fmt"
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
		Long:  `Initialize project. Add directories`,
		Args:  nil,
		Run: func(cmd *cobra.Command, args []string) {

			// 1. read config file
			// 2. create project dir
			// 3. create

			fmt.Println("This is the init command")

			//compound.Execute()

		},
	}
}
