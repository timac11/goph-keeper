package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timac11/goph-keeper/internal/client/model"
)

func BuildLoginCmd(executor func(model.LoginArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login user",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
