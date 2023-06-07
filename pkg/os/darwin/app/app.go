//go:build darwin
// +build darwin

package app

import (
	"github.com/adrianriobo/goax/pkg/os/darwin/api/appkit"
)

type appHandler struct {
	ref *appkit.NSRunningApplication
}

func (a *appHandler) Click(elementID string) error {
	return a.ref.Click(elementID)
}

func (a *appHandler) Check(elementID string) error {
	return a.ref.Check(elementID)
}

func GetApplication() (*appHandler, error) {
	ref, err := appkit.GetApplication()
	return &appHandler{ref: ref}, err
}
