//go:build windows

package ax

import (
	winax "github.com/adrianriobo/goax/pkg/os/windows/ax"
)

func osGetAXElement() (*AXElement, error) {
	rootAXElement, err := winax.GetForegroundRootAXElement()
	if err != nil {
		return nil, err
	}
	return GetAXElement(rootAXElement, nil)
}
