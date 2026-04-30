package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildUploadFileCmd(executor func(context.Context, model.UploadFileArgs) error) *cobra.Command {
	var commandArgs model.UploadFileArgs

	cmd := &cobra.Command{
		Use:   "upload-file",
		Short: "Upload file secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			if commandArgs.Path == "" {
				return errors.New("path is required")
			}

			return executor(cmd.Context(), commandArgs)
		},
	}

	cmd.Flags().StringVar(&commandArgs.Path, "path", "", "Path to uploaded file")
	cmd.Flags().StringVar(&commandArgs.Metadata, "metadata", "", "Secret metadata")

	return cmd
}
