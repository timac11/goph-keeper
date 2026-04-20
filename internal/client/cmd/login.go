package cmd

import (
	"github.com/spf13/cobra"
)

type LoginCmdArgs struct {
	Login    string
	Password string
}

func BuildLoginCmd(executor func(LoginCmdArgs) error) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login user",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: use executor
		},
	}
}
