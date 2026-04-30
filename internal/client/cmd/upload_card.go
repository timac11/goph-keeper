package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildUploadCardCmd(executor func(context.Context, model.UploadCardArgs) error) *cobra.Command {
	var commandArgs model.UploadCardArgs

	cmd := &cobra.Command{
		Use:   "upload-card",
		Short: "Upload card secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			if commandArgs.Name == "" {
				return errors.New("name is required")
			}
			if commandArgs.Number == "" {
				return errors.New("number is required")
			}
			if commandArgs.CVV == "" {
				return errors.New("cvv is required")
			}
			if commandArgs.UploadName == "" {
				return errors.New("upload name is required")
			}

			return executor(cmd.Context(), commandArgs)
		},
	}

	cmd.Flags().StringVar(&commandArgs.Name, "name", "", "Name of card owner")
	cmd.Flags().StringVar(&commandArgs.Number, "num", "", "Num of card")
	cmd.Flags().StringVar(&commandArgs.CVV, "cvv", "", "CVV of card")
	cmd.Flags().StringVar(&commandArgs.UploadName, "upload-name", "", "Upload name")
	cmd.Flags().StringVar(&commandArgs.Metadata, "metadata", "", "Secret metadata")

	return cmd
}
