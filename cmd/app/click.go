package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
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
	c.Flags().AddFlagSet(flagSet)
	c.MarkFlagRequired("element")
	return c
}

func click() error {
	a, err := app.LoadForefrontApp()
	if err != nil {
		return err
	}
	return a.ClickWithOrder(
		viper.GetString("element"),
		viper.GetString("element-type"),
		viper.IsSet("strict"),
		viper.GetInt("order"))
}
