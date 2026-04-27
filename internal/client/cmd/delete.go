package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildDeleteCmd(executor func(context.Context, model.DeleteArgs) error) *cobra.Command {
	var commandArgs model.DeleteArgs

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete data by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			if commandArgs.ID == "" {
				return errors.New("id is required")
			}

			return executor(cmd.Context(), commandArgs)
		},
	}

	cmd.Flags().StringVar(&commandArgs.ID, "id", "", "Id of secret to delete")

	return cmd
}
