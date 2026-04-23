package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildListCmd(executor func(model.ListArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List of uploaded data",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
