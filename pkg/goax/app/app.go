package app

import (
	"fmt"

	ax "github.com/adrianriobo/goax/pkg/goax/axapp"
	"github.com/adrianriobo/goax/pkg/goax/elements"
)

func Open(appPath string) error {
	return osOpen(appPath)
}

func LoadForefrontApp() (*app, error) {
	handler, err := osLoad()
	if err != nil {
		return nil, err
	}
	return &app{handler: handler}, nil
}

func (a *app) Click(element, elementType string, strict bool) error {
	return a.ClickWithOrder(element, elementType, strict, 0)
}

func (a *app) ClickWithOrder(element, elementType string, strict bool, order int) error {
	et, err := elements.GetElementType(elementType)
	if err != nil {
		return fmt.Errorf("Error running exists function: %v", err)
	}
	return a.handler.ClickWithOrder(element, et, strict, int8(order))
}

func (a *app) SetValue(element, elementType string, strict bool, value string) error {
	return a.SetValueWithOrder(element, elementType, strict, 0, value)
}

func (a *app) SetValueWithOrder(element, elementType string, strict bool, order int, value string) error {
	et, err := elements.GetElementType(elementType)
	if err != nil {
		return fmt.Errorf("Error running exists function: %v", err)
	}
	return a.handler.SetValueWithOrder(element, et, strict, int8(order), value)
}

// Check if an element exists within the app with the element id/value
func (a *app) Exists(element, elementType string, strict bool) (bool, error) {
	et, err := elements.GetElementType(elementType)
	if err != nil {
		return false, fmt.Errorf("Error running exists function: %v", err)
	}
	return a.handler.Exists(element, et, strict)
}

func (a *app) Print(idFilter string, strict bool) {
	a.handler.Print(idFilter, strict)
}

func osLoad() (appHandler, error) {
	//need to create an inspect a windows app
	return ax.GetAXApp()
}
