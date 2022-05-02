package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
		Args:  nil,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("This is the update command")

			compound.Execute()

		},
	}
}
