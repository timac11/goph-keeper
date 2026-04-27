package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildGetCmd(executor func(model.GetArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "get secret",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
