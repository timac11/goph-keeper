package cmd

import (
	"github.com/spf13/cobra"
)

type ListCmdArgs struct {}

func BuildListCmd(executor func(ListCmdArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List of uploaded data",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
