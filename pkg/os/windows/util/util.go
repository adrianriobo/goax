//go:build windows

package util

import (
	win32waf "github.com/adrianriobo/goax/pkg/os/windows/api/user-interface/windows-accesibility-features"
)

func Initialize() {
	// Initialize context
	win32waf.Initalize()
}

func Finalize() {
	// Finalize context
	win32waf.Finalize()
}
