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

func getSetValueCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "set-value",
		Short: "set a value",
		Long:  "set a value on a textbox or selectable element within the app",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return setValue()
		},
	}

	// Command flags
	flagSet := pflag.NewFlagSet("set-value", pflag.ExitOnError)
	flagSet.StringP("element", "e", "", "element id/value to be clicked")
	flagSet.StringP("element-type", "t", "", "element type to be clicked")
	flagSet.Bool("strict", false, "to force id match exactly with the element id")
	flagSet.Int("order", 0, "in case multiple elements with same id, we can specify the order of the element within the list of elements")
	flagSet.Bool("focus", false, "if focus flag is added it will add the value to the current focused textbox on the screen (if any)")
	flagSet.StringP("value", "v", "", "value to be set on the element")
	flagSet.Bool(record, false, recordDesc)
	flagSet.StringP(recordsPath, "", recordsPathDefault, recordsPathDesc)
	c.Flags().AddFlagSet(flagSet)
	c.MarkFlagRequired("value")
	c.MarkFlagsMutuallyExclusive("focus", "element")
	return c
}

func setValue() error {
	a, err := app.LoadForefrontApp()
	if err != nil {
		return err
	}
	delay.Delay(delay.LONG)
	if viper.IsSet(record) {
		if err := screenshot.CaptureScreen(viper.GetString(recordsPath), "setValueLoadForefrontApp"); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	if viper.IsSet("focus") {
		return a.SetValueOnFocus(viper.GetString("value"))
	}
	if err := a.SetValueWithOrder(
		viper.GetString("element"),
		viper.GetString("element-type"),
		viper.IsSet("strict"),
		viper.GetInt("order"),
		viper.GetString("value")); err != nil {
		return err
	}
	delay.Delay(delay.LONG)
	if viper.IsSet(record) {
		sfn := fmt.Sprintf("setValue-%s", viper.GetString("element"))
		if err := screenshot.CaptureScreen(viper.GetString(recordsPath), sfn); err != nil {
			logging.Errorf("error capturing the screenshot: %v", err)
		}
	}
	return nil
}
