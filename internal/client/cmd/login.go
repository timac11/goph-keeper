package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildLoginCmd(executor func(context.Context, model.LoginArgs) error) *cobra.Command {
	var commandArgs model.LoginArgs

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login user",
		RunE: func(cmd *cobra.Command, args []string) error {
			if commandArgs.Login == "" {
				return errors.New("login is required")
			}
			if commandArgs.Password == "" {
				return errors.New("password is required")
			}

			err := executor(cmd.Context(), commandArgs)
			if err != nil {
				return err
			}

			cmd.Println("Authorization completed!")
			return nil
		},
	}

	cmd.Flags().StringVar(&commandArgs.Login, "login", "", "User login")
	cmd.Flags().StringVar(&commandArgs.Password, "password", "", "User password")

	return cmd
}
