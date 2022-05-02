package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CommandRepo struct {
	Root *cobra.Command
}

func NewCommandRepo() *CommandRepo {
	var rootCmd = &cobra.Command{
		Use:   "gitv",
		Short: "Gitlab Variables",
		Long:  `Gitlab Variables`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Nothing to see here.")
		},
	}
	return &CommandRepo{rootCmd}
}
