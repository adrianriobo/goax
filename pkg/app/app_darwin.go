package app

import (
	"os/exec"
)

func osOpenApp(appPath string) error {
	cmd := exec.Command("open", appPath)
	return cmd.Start()
}
