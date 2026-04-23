package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildRegisterCmd(executor func(model.RegisterArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Register user",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
