//go:build darwin
// +build darwin

package ax

import (
	darwinax "github.com/adrianriobo/goax/pkg/os/darwin/ax"
)

func osGetAXElement() (*AXElement, error) {
	rootAXElement, err := darwinax.GetForegroundRootAXElement()
	if err != nil {
		return nil, err
	}
	return GetAXElement(rootAXElement, nil)
}
