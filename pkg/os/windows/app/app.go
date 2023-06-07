//go:build windows
// +build windows

package app

import (
	"github.com/adrianriobo/goax/pkg/os/windows/app/ux"
)

type appHandler struct {
	ref *ux.UXElement
}

func (a *appHandler) Click(elementID string) (err error) {
	err = a.ref.PressElementByID(elementID)
	// TODO review this here if it is used as library
	ux.Finalize()
	return
}

func (a *appHandler) Check(elementID string) (err error) {
	err = a.ref.PressElementByID(elementID)
	ux.Finalize()
	return
}

func GetApplication() (*appHandler, error) {
	// Move this to cmd or functions
	ux.Initialize()
	parentRef, err := ux.GetActiveElement("Podman Desktop", ux.PANE)
	if err != nil {
		return nil, err
	}
	ref, err := ux.GetUXElementRef(parentRef.Ref, nil)
	if err != nil {
		return nil, err
	}
	// Print for debugging purposes
	// ref.Print()
	return &appHandler{ref: ref}, nil
}
