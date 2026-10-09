package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildGetCmd(executor func(context.Context, model.GetArgs) (string, error)) *cobra.Command {
	var commandArgs model.GetArgs

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			if commandArgs.ID == "" {
				return errors.New("secret id is required")
			}

			_, err := executor(cmd.Context(), commandArgs)

			// TODO: print info about secret

			return err
		},
	}

	cmd.Flags().StringVar(&commandArgs.ID, "id", "", "Secret ID")

	return cmd
}
