package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrianriobo/goax/cmd/app"
	"github.com/adrianriobo/goax/pkg/util"
	"github.com/adrianriobo/goax/pkg/util/logging"
	"github.com/spf13/cobra"
	"k8s.io/utils/exec"
)

const (
	commandName      = "goax"
	descriptionShort = "ux management through AX API"
	descriptionLong  = "ux management through AX API"

	defaultErrorExitCode = 1
)

var (
	rootCmd *cobra.Command

	baseDir = filepath.Join(util.GetHomeDir(), ".goax")
	logFile = "goax.log"
)

func main() {
	attachMiddleware([]string{}, rootCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		runPostrun()
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		var e exec.CodeExitError
		if errors.As(err, &e) {
			os.Exit(e.ExitStatus())
		} else {
			os.Exit(defaultErrorExitCode)
		}
	}
	runPostrun()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   commandName,
		Short: descriptionShort,
		Long:  descriptionLong,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runPrerun(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			runRoot()
			_ = cmd.Help()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// Subcommands
	rootCmd.AddCommand(
		app.GetCmd())
}

func runPrerun(cmd *cobra.Command) error {
	logging.InitLogrus(logging.LogLevel, baseDir, logFile)
	return nil
}

func runRoot() {
	fmt.Println("No command given")
}

func attachMiddleware(names []string, cmd *cobra.Command) {
	if cmd.HasSubCommands() {
		for _, command := range cmd.Commands() {
			attachMiddleware(append(names, cmd.Name()), command)
		}
	} else if cmd.RunE != nil {
		fullCmd := strings.Join(append(names, cmd.Name()), " ")
		src := cmd.RunE
		cmd.RunE = executeWithLogging(fullCmd, src)
	}
}

func executeWithLogging(fullCmd string, input func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		logging.Debugf("Running '%s'", fullCmd)
		return input(cmd, args)
	}
}

func runPostrun() {
	logging.CloseLogging()
}
