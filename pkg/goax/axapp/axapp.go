package ax

import (
	"fmt"

	"github.com/adrianriobo/goax/pkg/goax/elements"
)

// AXApp represent an app which is accessible through
// an hierarchy of accessible elements through accessible APIS
// It implements appHandler
type AXApp struct {
	ref *AXElement
}

var (
	clickableElements = []elements.ElementType{
		elements.BUTTON, elements.HYPERLINK, elements.TEXT}
	checkableElements = []elements.ElementType{
		elements.CHECKBOX}
	setableElements = []elements.ElementType{
		elements.TEXT, elements.EDIT}
)

// Get the frontground app from the system
// and the hirarchy of accessible elements
func GetAXApp() (*AXApp, error) {
	ref, err := osGetAXElement()
	if err != nil {
		return nil, err
	}
	return &AXApp{
		ref: ref,
	}, nil
}

// Exists(element string, elementType *elements.ElementType, strict bool) error
// // Click on a button within the app
// Click(element string, elementType *elements.ElementType, strict bool) error
// ClickWithOrder(element string, elementType *elements.ElementType, strict bool, order int8) error
// // Set the value on an element within the app
// SetValue(element string, elementType *elements.ElementType, strict bool, value string) error
// SetValueWithOrder(element string, elementType *elements.ElementType, strict bool, order int8, value string) error
// SetValueOnFocus(value string) error

func (a *AXApp) Print(elementFilter string, strict bool) {
	a.ref.Print(true, elementFilter, strict)
}

func (a *AXApp) Exists(element string, elementType *elements.ElementType, strict bool) (bool, error) {
	// a.ref.Print(true, elementFilter, strict)
	elementTypes := clickableElements
	if elementType != nil {
		elementTypes = []elements.ElementType{*elementType}
	}
	e, err := a.findElement(element, strict, elementTypes, 0)
	return e != nil, err
}

func (a *AXApp) Click(element string, elementType *elements.ElementType, strict bool) error {
	return a.ClickWithOrder(element, elementType, strict, 0)
}

// func (a *AXApp) Check(id string, strict bool) error {
// 	return a.press(id, strict, checkableElements, 0)
// }

func (a *AXApp) ClickWithOrder(element string, elementType *elements.ElementType, strict bool, order int8) error {
	// return a.press(id, strict, clickableElements, order)
	elementTypes := clickableElements
	if elementType != nil {
		elementTypes = []elements.ElementType{*elementType}
	}
	return a.press(element, strict, elementTypes, order)
}

// func (a *AXApp) CheckWithOrder(id string, strict bool, order int8) error {
// 	return a.press(id, strict, checkableElements, order)
// }

func (a *AXApp) SetValue(element string, elementType *elements.ElementType, strict bool, value string) error {
	return a.SetValueWithOrder(element, elementType, strict, 0, value)
}

func (a *AXApp) SetValueWithOrder(element string, elementType *elements.ElementType, strict bool, order int8, value string) error {
	elementTypes := setableElements
	if elementType != nil {
		elementTypes = []elements.ElementType{*elementType}
	}
	return a.setValue(element, strict, elementTypes, order, value)
}

func (a *AXApp) SetValueOnFocus(value string) error {
	return a.ref.Ref.SetValueOnFocus(value)
}

func (a *AXApp) press(id string, strict bool, elementTypes []elements.ElementType, order int8) error {
	element, err := a.findElement(id, strict, elementTypes, order)
	if err != nil {
		return err
	}
	err = element.Ref.Press()
	if err != nil {
		return fmt.Errorf("error pressing element %s: %v", id, err)
	}
	return nil
}

func (a *AXApp) setValue(id string, strict bool, elementTypes []elements.ElementType, order int8, value string) error {
	element, err := a.findElement(id, strict, elementTypes, order)
	if err != nil {
		return err
	}
	err = element.Ref.SetValue(value)
	if err != nil {
		return fmt.Errorf("error setting value to element %s: %v", id, err)
	}
	return nil
}

func (a *AXApp) findElement(id string, strict bool, elementTypes []elements.ElementType, order int8) (*AXElement, error) {
	elements := a.ref.FindElements(id, elementTypes, strict)
	l := len(elements)
	if l == 0 {
		return nil, fmt.Errorf("can not find any element with id %s", id)
	}
	if l > 1 && order == 0 {
		return nil, fmt.Errorf("there are %d elements matching the id, please specifiy the order", l)
	}
	// If there is only one element matching directly press on it
	if order > 0 {
		order = order - 1
	}
	return elements[order], nil
}
