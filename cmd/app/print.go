package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
)

func getPrintCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "print",
		Short: "print app ax elements",
		Long:  "print app ax elements",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return print()
		},
	}

	// // Command flags
	flagSet := pflag.NewFlagSet("print", pflag.ExitOnError)
	flagSet.StringP("element", "e", "", "element id/value to be filtered")
	flagSet.Bool("strict", false, "to force id match exactly with the element id/value")
	flagSet.StringP("type", "t", "", "type of element. This is OS dependant. For windows it is the type of element (i.e Window) for mac is the bundle id")
	flagSet.StringP("title", "", "", "title of the main window")
	c.Flags().AddFlagSet(flagSet)

	return c
}

func print() error {
	var a *app.App
	var err error
	aType := viper.GetString("type")
	aTitle := viper.GetString("title")
	if len(aType) > 0 && len(aTitle) > 0 {
		a, err = app.LoadAppByTypeAndTitle(
			aType,
			aTitle,
		)
	} else {
		a, err = app.LoadForefrontApp()
	}
	if err != nil {
		return err
	}
	a.Print(
		viper.GetString("element"),
		viper.IsSet("strict"))
	return nil
}
