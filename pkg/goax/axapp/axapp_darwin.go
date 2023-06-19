//go:build darwin
// +build darwin

package ax

import (
	darwinax "github.com/adrianriobo/goax/pkg/os/darwin/ax"
	"github.com/adrianriobo/goax/pkg/util/delay"
)

func osGetAXElement() (*AXElement, error) {
	rootAXElement, err := darwinax.GetForegroundRootAXElement()
	if err != nil {
		return nil, err
	}
	delay.Delay(delay.MEDIUM)
	return GetAXElement(rootAXElement, nil)
}
