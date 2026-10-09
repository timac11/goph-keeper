package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildRegisterCmd(executor func(context.Context, model.RegisterArgs) error) *cobra.Command {
	var commandArgs model.RegisterArgs

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register user",
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

			cmd.Println("Registration completed!")
			return nil
		},
	}

	cmd.Flags().StringVar(&commandArgs.Login, "login", "", "User login")
	cmd.Flags().StringVar(&commandArgs.Password, "password", "", "User password")

	return cmd
}
