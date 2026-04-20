package cmd

import (
	"github.com/spf13/cobra"
)

type DeleteCmdArgs struct{}

func BuildDeleteCmd(executor func(DeleteCmdArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete data by id",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
