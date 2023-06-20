package api

import "github.com/adrianriobo/goax/pkg/goax/elements"

type AppHandler interface {
	Reload() (AppHandler, error)
	// Prints the accessible elements within the app
	Print(elementFilter string, strict bool)
	// Check if exists an element with the id
	Exists(element string, elementType *elements.ElementType, strict bool) (bool, error)
	// Click on a button within the app
	Click(element string, elementType *elements.ElementType, strict bool) error
	ClickWithOrder(element string, elementType *elements.ElementType, strict bool, order int8) error
	// Set the value on an element within the app
	SetValue(element string, elementType *elements.ElementType, strict bool, value string) error
	SetValueWithOrder(element string, elementType *elements.ElementType, strict bool, order int8, value string) error
	SetValueOnFocus(value string) error
}
