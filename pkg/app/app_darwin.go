//go:build darwin
// +build darwin

package app

import (
	"os/exec"

	darwinApp "github.com/adrianriobo/goax/pkg/os/darwin/app"
)

func osOpen(appPath string) error {
	cmd := exec.Command("open", appPath)
	return cmd.Start()
}

func osLoad() (appHandler, error) {
	return darwinApp.GetApplication()
}
