//go:build darwin
// +build darwin

package app

import (
	"os/exec"
)

func osOpen(appPath string) error {
	cmd := exec.Command("open", appPath)
	return cmd.Start()
}
