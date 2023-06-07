//go:build darwin
// +build darwin

package app

import (
	"github.com/adrianriobo/goax/pkg/os/darwin/api/appkit"
)

type darwinAppHandler struct {
	ref *appkit.NSRunningApplication
}

func (h darwinAppHandler) Click(elementID string) error {
	return h.ref.Click(elementID)
}

func (h darwinAppHandler) Check(elementID string) error {
	return h.ref.Check(elementID)
}

func GetApplication() (*darwinAppHandler, error) {
	ref, err := appkit.GetApplication()
	return &darwinAppHandler{ref: ref}, err
}
