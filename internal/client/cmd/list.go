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
			secrets, err := executor(cmd.Context(), commandArgs)

			if err != nil {
				return err
			}

			cmd.Printf("Available %d secrets:", len(*secrets))

			for _, secret := range *secrets {
				cmd.Println("%s, %s, %s", secret.ID, secret.Name, secret.Type)
			}

			return nil
		},
	}

	return cmd
}
