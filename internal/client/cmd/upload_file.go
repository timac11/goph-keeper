package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildUploadFileCmd(executor func(model.UploadFileArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "upload-file",
		Short: "Upload file",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
