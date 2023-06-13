package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/goax/pkg/goax/app"
)

func getExistsCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "exists",
		Short: "check if is there an element within the application with the id/value",
		Long:  "check if is there an element within the application with the id/value",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return check()
		},
	}

	// Command flags
	flagSet := pflag.NewFlagSet("exists", pflag.ExitOnError)
	flagSet.StringP("element", "e", "", "element id/value to be checked")
	flagSet.StringP("element-type", "t", "", "element type to be checked")
	flagSet.Bool("strict", false, "if set id/value should match exactly")
	c.Flags().AddFlagSet(flagSet)
	c.MarkFlagRequired("element")

	return c
}

func check() error {
	a, err := app.LoadForefrontApp()
	if err != nil {
		return err
	}
	_, err = a.Exists(viper.GetString("element"), viper.GetString("element-type"), viper.IsSet("strict"))
	return err
}
