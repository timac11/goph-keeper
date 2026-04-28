package cmd

import (
	"context"

	"github.com/spf13/cobra"

	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

func BuildListCmd(executor func(context.Context, clientModel.ListArgs) (*[]model.SecretInfoDto, error)) *cobra.Command {
	var commandArgs clientModel.ListArgs

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List of uploaded data",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := executor(cmd.Context(), commandArgs)

			if err != nil {
				return err
			}

			// TODO: print list

			return nil
		},
	}

	return cmd
}
