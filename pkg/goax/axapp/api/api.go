package api

import (
	"github.com/adrianriobo/goax/pkg/goax/elements"
)

// Interface which defines the operations which
// should be executable for each accessible element
// within the app
type OSAXElement interface {
	Press() error
	SetValue(value string) error
	SetValueOnFocus(value string) error
	GetAllChildren() ([]OSAXElement, error)
	GetType() (*elements.ElementType, error)
	GetID() (string, error)
	GetValue() (string, error)
}
