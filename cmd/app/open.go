package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
)

func getOpenCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "open",
		Short: "open an application",
		Long:  "open an application based on its path",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return open()
		},
	}

	flagSet := pflag.NewFlagSet("app", pflag.ExitOnError)
	flagSet.StringP(appPath, "p", "", "path for the application to be handle")
	c.Flags().AddFlagSet(flagSet)

	return c
}

func open() error {
	return app.Open(viper.GetString(appPath))
}
