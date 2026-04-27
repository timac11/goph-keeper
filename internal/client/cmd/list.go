package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildListCmd(executor func(context.Context, model.ListArgs) error) *cobra.Command {
	var commandArgs model.ListArgs

	return &cobra.Command{
		Use:   "list",
		Short: "List of uploaded data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return executor(cmd.Context(), commandArgs)
		},
	}
}
