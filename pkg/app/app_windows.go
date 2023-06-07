package app

import (
	"os/exec"
	"time"

	windowsApp "github.com/adrianriobo/goax/pkg/os/windows/app"
)

func osOpen(appPath string) error {
	cmd := exec.Command(appPath)
	if err := cmd.Start(); err != nil {
		return err
	}
	// delay to get window as active
	time.Sleep(1 * time.Second)
	return nil
}

func osLoad() (appHandler, error) {
	//need to create an inspect a windows app
	return windowsApp.GetApplication()
}
