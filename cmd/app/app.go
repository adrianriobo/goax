package app

import (
	"github.com/spf13/cobra"
)

const (
	appPath         = "app-path"
	record          = "record"
	recordDesc      = "capture screenshot"
	recordsPath     = "records-path"
	recordsPathDesc = "path to store the screeshots"
)

var (
	recordsPathDefault = "."
)

func GetCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "app",
		Short: "handle app ux",
		Long:  "handle app ux"}

	// Subcommands
	c.AddCommand(
		getOpenCmd(),
		getClickCmd(),
		getExistsCmd(),
		getPrintCmd(),
		getSetValueCmd())

	return c
}
