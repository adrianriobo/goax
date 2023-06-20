//go:build windows

package ax

import (
	axAPI "github.com/adrianriobo/goax/pkg/goax/axapp/api"
	winax "github.com/adrianriobo/goax/pkg/os/windows/ax"
)

func osGetAXElement(element *AXElement) (*AXElement, error) {
	var rootAXElement axAPI.OSAXElement
	var err error
	if element != nil {
		rootAXElement = element.Ref
	} else {
		rootAXElement, err = winax.GetForegroundRootAXElement()
		if err != nil {
			return nil, err
		}
	}
	return GetAXElement(rootAXElement, nil)
}
