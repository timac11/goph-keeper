package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildUploadCardCmd(executor func(model.UploadCardArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "upload-card",
		Short: "Upload card info",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
