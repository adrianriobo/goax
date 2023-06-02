package app

import (
	"os/exec"
	"time"
)

func osOpenApp(appPath string) error {
	cmd := exec.Command(appPath)
	if err := cmd.Start(); err != nil {
		return err
	}
	// delay to get window as active
	time.Sleep(1 * time.Second)
	return nil
}
