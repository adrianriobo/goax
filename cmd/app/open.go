package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/goax/pkg/util/logging"
	"github.com/adrianriobo/goax/pkg/util/screenshot"
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
	flagSet.Bool(record, false, recordDesc)
	flagSet.StringP(recordsPath, "", recordsPathDefault, recordsPathDesc)
	c.Flags().AddFlagSet(flagSet)
	return c
}

func open() error {
	if err := app.Open(
		viper.GetString(appPath)); err != nil {
		return err
	}
	// We open remotely so we wait for a bit
	delay.Delay(delay.LONG)
	if viper.IsSet(record) {
		if err := screenshot.CaptureScreen(viper.GetString(recordsPath), "openApp"); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	return nil
}
