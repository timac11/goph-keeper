package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildDeleteCmd(executor func(model.DeleteArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete data by id",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
