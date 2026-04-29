package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildUploadAuthCmd(executor func(context.Context, model.UploadAuthArgs) error) *cobra.Command {
	var commandArgs model.UploadAuthArgs

	cmd := &cobra.Command{
		Use:   "upload-auth",
		Short: "Upload auth info",
		RunE: func(cmd *cobra.Command, args []string) error {
			if commandArgs.Login == "" {
				return errors.New("login is required")
			}
			if commandArgs.Password == "" {
				return errors.New("password is required")
			}
			if commandArgs.UploadName == "" {
				return errors.New("upload name is required")
			}

			return executor(cmd.Context(), commandArgs)
		},
	}

	cmd.Flags().StringVar(&commandArgs.Login, "login", "", "Account login")
	cmd.Flags().StringVar(&commandArgs.Password, "password", "", "Account password")
	cmd.Flags().StringVar(&commandArgs.UploadName, "upload", "", "Upload name")

	return cmd
}
