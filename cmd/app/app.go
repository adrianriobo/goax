package app

import (
	"github.com/spf13/cobra"
)

const (
	appPath = "app-path"

	onlyOpen    bool = false
	openAndLoad bool = true
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
		getCheckCmd())

	return c
}
