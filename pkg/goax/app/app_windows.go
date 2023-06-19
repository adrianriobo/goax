//go:build windows
// +build windows

package app

import (
	"fmt"
	"github.com/adrianriobo/goax/pkg/os/windows/powershell"
	"github.com/adrianriobo/goax/pkg/util/delay"
)

func osOpen(appPath string) error {
	// We will run commands through ps
	ps := powershell.New()
	// Minize everything
	minimizeall := "-c (New-Object -ComObject \"Shell.Application\").minimizeall()"
	ps.Execute(minimizeall)
	sendEsc := "(New-Object -ComObject wscript.shell).SendKeys(\"{ESC}\")"
	ps.Execute(sendEsc)
	// To ensure window get foreground on windows 11 we need to open if from a local process
	openApp := fmt.Sprintf("Start-Process powershell.exe -WindowStyle Hidden -ArgumentList \"%s\"", appPath)
	ps.Execute(openApp)
	delay.Delay(delay.XLONG)
	// This is a hack to pre load ux elements on OS
	_, err := osLoad()
	return err
}
