package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
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
	flagSet.StringP("element", "", "", "element id/value to be clicked")
	flagSet.StringP("element-type", "t", "", "element type to be clicked")
	flagSet.Bool("strict", false, "to force id match exactly with the element id")
	flagSet.Int("order", 0, "in case multiple elements with same id, we can specify the order of the element within the list of elements")
	flagSet.StringP("value", "v", "", "value to be set on the element")
	c.Flags().AddFlagSet(flagSet)
	c.MarkFlagRequired("element")
	return c
}

func setValue() error {
	a, err := app.LoadForefrontApp()
	if err != nil {
		return err
	}
	return a.SetValueWithOrder(
		viper.GetString("element"),
		viper.GetString("element-type"),
		viper.IsSet("strict"),
		viper.GetInt("order"),
		viper.GetString("value"))
}
