package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/app"
)

func getCheckCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "check",
		Short: "check on an application element",
		Long:  "check on an application element",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return check()
		},
	}

	// Command flags
	flagSet := pflag.NewFlagSet("click", pflag.ExitOnError)
	flagSet.StringP(appPath, "p", "", "path for the application to be handle")
	flagSet.StringP("id", "", "", "id for the element to be checked")
	c.Flags().AddFlagSet(flagSet)

	return c
}

func check() error {
	a := app.New(viper.GetString(appPath))
	if err := a.Open(openAndLoad); err != nil {
		return err
	}
	return a.Check(viper.GetString("id"))
}
