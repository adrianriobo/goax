//go:build darwin
// +build darwin

package ax

import (
	axAPI "github.com/adrianriobo/goax/pkg/goax/axapp/api"
	darwinax "github.com/adrianriobo/goax/pkg/os/darwin/ax"
	"github.com/adrianriobo/goax/pkg/util/delay"
)

func osGetAXElement(element *AXElement) (*AXElement, error) {
	var rootAXElement axAPI.OSAXElement
	var err error
	if element != nil {
		rootAXElement = element.Ref
	} else {
		rootAXElement, err = darwinax.GetForegroundRootAXElement()
		if err != nil {
			return nil, err
		}
	}
	delay.Delay(delay.MEDIUM)
	return GetAXElement(rootAXElement, nil)
}
