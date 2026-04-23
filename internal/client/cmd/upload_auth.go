package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildUploadAuthCmd(executor func(model.UploadAuthArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "upload-auth",
		Short: "Upload auth info",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
