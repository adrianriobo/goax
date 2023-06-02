package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/app"
)

const (
	appPath = "path"
)

func GetOpenCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "open",
		Short: "Open an application",
		Long:  "Open an application based on its path",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return open()
		},
	}

	// Command flags
	flagSet := pflag.NewFlagSet("open", pflag.ExitOnError)
	flagSet.StringP(appPath, "p", "", "path for the application to be open")
	c.Flags().AddFlagSet(flagSet)

	return c
}

func open() error {
	a := app.New(viper.GetString(appPath))
	return a.Open()
}
