package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
	"github.com/adrianriobo/goax/pkg/util/delay"
	"github.com/adrianriobo/goax/pkg/util/logging"
	"github.com/adrianriobo/goax/pkg/util/screenshot"
)

func getClickCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "click",
		Short: "click on an application element",
		Long:  "click on an application element",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return click()
		},
	}

	// Command flags
	flagSet := pflag.NewFlagSet("click", pflag.ExitOnError)
	flagSet.StringP("element", "e", "", "element id/value to be clicked")
	flagSet.StringP("element-type", "t", "", "element type to be clicked")
	flagSet.Bool("strict", false, "to force id match exactly with the element id")
	flagSet.Int("order", 0, "in case multiple elements with same id, we can specify the order of the element within the list of elements")
	flagSet.Bool(record, false, recordDesc)
	flagSet.StringP(recordsPath, "", recordsPathDefault, recordsPathDesc)
	c.Flags().AddFlagSet(flagSet)
	c.MarkFlagRequired("element")
	return c
}

func click() error {
	a, err := app.LoadForefrontApp()
	if err != nil {
		return err
	}
	delay.Delay(delay.LONG)
	if viper.IsSet(record) {
		if err := screenshot.CaptureScreen(viper.GetString(recordsPath), "clickLoadForefrontApp"); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	if err := a.ClickWithOrder(
		viper.GetString("element"),
		viper.GetString("element-type"),
		viper.IsSet("strict"),
		viper.GetInt("order")); err != nil {
		return err
	}
	delay.Delay(delay.LONG)
	if viper.IsSet(record) {
		sfn := fmt.Sprintf("click-%s", viper.GetString("element"))
		if err := screenshot.CaptureScreen(viper.GetString(recordsPath), sfn); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	return nil
}
